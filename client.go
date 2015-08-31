package bbs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudfoundry-incubator/bbs/events"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/cloudfoundry-incubator/cf_http"
	"github.com/gogo/protobuf/proto"
	"github.com/tedsuo/rata"
	"github.com/vito/go-sse/sse"
)

const (
	ContentTypeHeader    = "Content-Type"
	XCfRouterErrorHeader = "X-Cf-Routererror"
	ProtoContentType     = "application/x-protobuf"
	KeepContainer        = true
	DeleteContainer      = false
)

//go:generate counterfeiter -o fake_bbs/fake_client.go . Client

type Client interface {
	Domains() ([]string, *models.Error)
	UpsertDomain(domain string, ttl time.Duration) *models.Error

	ActualLRPGroups(models.ActualLRPFilter) ([]*models.ActualLRPGroup, *models.Error)
	ActualLRPGroupsByProcessGuid(processGuid string) ([]*models.ActualLRPGroup, *models.Error)
	ActualLRPGroupByProcessGuidAndIndex(processGuid string, index int) (*models.ActualLRPGroup, *models.Error)

	ClaimActualLRP(processGuid string, index int, instanceKey *models.ActualLRPInstanceKey) *models.Error
	StartActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, netInfo *models.ActualLRPNetInfo) *models.Error
	CrashActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, errorMessage string) *models.Error
	FailActualLRP(key *models.ActualLRPKey, errorMessage string) *models.Error
	RemoveActualLRP(processGuid string, index int) *models.Error
	RetireActualLRP(key *models.ActualLRPKey) *models.Error

	EvacuateClaimedActualLRP(*models.ActualLRPKey, *models.ActualLRPInstanceKey) (bool, *models.Error)
	EvacuateRunningActualLRP(*models.ActualLRPKey, *models.ActualLRPInstanceKey, *models.ActualLRPNetInfo, uint64) (bool, *models.Error)
	EvacuateStoppedActualLRP(*models.ActualLRPKey, *models.ActualLRPInstanceKey) (bool, *models.Error)
	EvacuateCrashedActualLRP(*models.ActualLRPKey, *models.ActualLRPInstanceKey, string) (bool, *models.Error)
	RemoveEvacuatingActualLRP(*models.ActualLRPKey, *models.ActualLRPInstanceKey) *models.Error

	DesiredLRPs(models.DesiredLRPFilter) ([]*models.DesiredLRP, *models.Error)
	DesiredLRPByProcessGuid(processGuid string) (*models.DesiredLRP, *models.Error)

	DesireLRP(*models.DesiredLRP) *models.Error
	UpdateDesiredLRP(processGuid string, update *models.DesiredLRPUpdate) *models.Error
	RemoveDesiredLRP(processGuid string) *models.Error

	ConvergeLRPs() *models.Error

	// Public Task Methods
	Tasks() ([]*models.Task, *models.Error)
	TasksByDomain(domain string) ([]*models.Task, *models.Error)
	TasksByCellID(cellId string) ([]*models.Task, *models.Error)
	TaskByGuid(guid string) (*models.Task, *models.Error)

	DesireTask(guid, domain string, def *models.TaskDefinition) *models.Error
	CancelTask(taskGuid string) *models.Error
	FailTask(taskGuid, failureReason string) *models.Error
	CompleteTask(taskGuid, cellId string, failed bool, failureReason, result string) *models.Error
	ResolvingTask(taskGuid string) *models.Error
	DeleteTask(taskGuid string) *models.Error

	ConvergeTasks(kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration time.Duration) *models.Error

	SubscribeToEvents() (events.EventSource, *models.Error)

	// Internal Task Methods
	StartTask(taskGuid string, cellID string) (bool, *models.Error)
}

func NewClient(url string) Client {
	return &client{
		httpClient:          cf_http.NewClient(),
		streamingHTTPClient: cf_http.NewStreamingClient(),

		reqGen: rata.NewRequestGenerator(url, Routes),
	}
}

type client struct {
	httpClient          *http.Client
	streamingHTTPClient *http.Client

	reqGen *rata.RequestGenerator
}

func (c *client) Domains() ([]string, *models.Error) {
	response := models.DomainsResponse{}
	err := c.doRequest(DomainsRoute, nil, nil, nil, &response)
	if err != nil {
		return nil, err
	}
	return response.Domains, response.Error
}

func (c *client) UpsertDomain(domain string, ttl time.Duration) *models.Error {
	request := models.UpsertDomainRequest{
		Domain: domain,
		Ttl:    uint32(ttl.Seconds()),
	}
	response := models.UpsertDomainResponse{}
	err := c.doRequest(UpsertDomainRoute, nil, nil, &request, &response)
	if err != nil {
		return err
	}
	return response.Error
}

func (c *client) ActualLRPGroups(filter models.ActualLRPFilter) ([]*models.ActualLRPGroup, *models.Error) {
	request := models.ActualLRPGroupsRequest{
		Domain: filter.Domain,
		CellId: filter.CellID,
	}
	response := models.ActualLRPGroupsResponse{}
	err := c.doRequest(ActualLRPGroupsRoute, nil, nil, &request, &response)
	if err != nil {
		return nil, err
	}

	return response.ActualLrpGroups, response.Error
}

func (c *client) ActualLRPGroupsByProcessGuid(processGuid string) ([]*models.ActualLRPGroup, *models.Error) {
	request := models.ActualLRPGroupsByProcessGuidRequest{
		ProcessGuid: processGuid,
	}
	response := models.ActualLRPGroupsResponse{}
	err := c.doRequest(ActualLRPGroupsByProcessGuidRoute, nil, nil, &request, &response)
	if err != nil {
		return nil, err
	}

	return response.ActualLrpGroups, response.Error
}

func (c *client) ActualLRPGroupByProcessGuidAndIndex(processGuid string, index int) (*models.ActualLRPGroup, *models.Error) {
	request := models.ActualLRPGroupByProcessGuidAndIndexRequest{
		ProcessGuid: processGuid,
		Index:       int32(index),
	}
	response := models.ActualLRPGroupResponse{}
	err := c.doRequest(ActualLRPGroupByProcessGuidAndIndexRoute, nil, nil, &request, &response)
	if err != nil {
		return nil, err
	}

	return response.ActualLrpGroup, response.Error
}

func (c *client) ClaimActualLRP(processGuid string, index int, instanceKey *models.ActualLRPInstanceKey) *models.Error {
	request := models.ClaimActualLRPRequest{
		ProcessGuid:          processGuid,
		Index:                int32(index),
		ActualLrpInstanceKey: instanceKey,
	}
	response := models.ActualLRPLifecycleResponse{}
	err := c.doRequest(ClaimActualLRPRoute, nil, nil, &request, &response)
	if err != nil {
		return err
	}
	return response.Error
}

func (c *client) StartActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, netInfo *models.ActualLRPNetInfo) *models.Error {
	request := models.StartActualLRPRequest{
		ActualLrpKey:         key,
		ActualLrpInstanceKey: instanceKey,
		ActualLrpNetInfo:     netInfo,
	}
	response := models.ActualLRPLifecycleResponse{}
	err := c.doRequest(StartActualLRPRoute, nil, nil, &request, &response)
	if err != nil {
		return err

	}
	return response.Error
}

func (c *client) CrashActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, errorMessage string) *models.Error {
	request := models.CrashActualLRPRequest{
		ActualLrpKey:         key,
		ActualLrpInstanceKey: instanceKey,
		ErrorMessage:         errorMessage,
	}
	response := models.ActualLRPLifecycleResponse{}
	err := c.doRequest(CrashActualLRPRoute, nil, nil, &request, &response)
	if err != nil {
		return err

	}
	return response.Error
}

func (c *client) FailActualLRP(key *models.ActualLRPKey, errorMessage string) *models.Error {
	request := models.FailActualLRPRequest{
		ActualLrpKey: key,
		ErrorMessage: errorMessage,
	}
	response := models.ActualLRPLifecycleResponse{}
	err := c.doRequest(FailActualLRPRoute, nil, nil, &request, &response)
	if err != nil {
		return err

	}
	return response.Error
}

func (c *client) RetireActualLRP(key *models.ActualLRPKey) *models.Error {
	request := models.RetireActualLRPRequest{
		ActualLrpKey: key,
	}
	response := models.ActualLRPLifecycleResponse{}
	err := c.doRequest(RetireActualLRPRoute, nil, nil, &request, &response)
	if err != nil {
		return err

	}
	return response.Error
}

func (c *client) RemoveActualLRP(processGuid string, index int) *models.Error {
	request := models.RemoveActualLRPRequest{
		ProcessGuid: processGuid,
		Index:       int32(index),
	}
	response := models.ActualLRPLifecycleResponse{}
	err := c.doRequest(RemoveActualLRPRoute, nil, nil, &request, &response)
	if err != nil {
		return err
	}
	return response.Error
}

func (c *client) EvacuateClaimedActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey) (bool, *models.Error) {
	return c.doEvacRequest(EvacuateClaimedActualLRPRoute, KeepContainer, &models.EvacuateClaimedActualLRPRequest{
		ActualLrpKey:         key,
		ActualLrpInstanceKey: instanceKey,
	})
}

func (c *client) EvacuateCrashedActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, errorMessage string) (bool, *models.Error) {
	return c.doEvacRequest(EvacuateCrashedActualLRPRoute, DeleteContainer, &models.EvacuateCrashedActualLRPRequest{
		ActualLrpKey:         key,
		ActualLrpInstanceKey: instanceKey,
		ErrorMessage:         errorMessage,
	})
}

func (c *client) EvacuateStoppedActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey) (bool, *models.Error) {
	return c.doEvacRequest(EvacuateStoppedActualLRPRoute, DeleteContainer, &models.EvacuateStoppedActualLRPRequest{
		ActualLrpKey:         key,
		ActualLrpInstanceKey: instanceKey,
	})
}

func (c *client) EvacuateRunningActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey, netInfo *models.ActualLRPNetInfo, ttl uint64) (bool, *models.Error) {
	return c.doEvacRequest(EvacuateRunningActualLRPRoute, KeepContainer, &models.EvacuateRunningActualLRPRequest{
		ActualLrpKey:         key,
		ActualLrpInstanceKey: instanceKey,
		ActualLrpNetInfo:     netInfo,
		Ttl:                  ttl,
	})
}

func (c *client) RemoveEvacuatingActualLRP(key *models.ActualLRPKey, instanceKey *models.ActualLRPInstanceKey) *models.Error {
	request := models.RemoveEvacuatingActualLRPRequest{
		ActualLrpKey:         key,
		ActualLrpInstanceKey: instanceKey,
	}

	response := models.RemoveEvacuatingActualLRPResponse{}
	err := c.doRequest(RemoveEvacuatingActualLRPRoute, nil, nil, &request, &response)
	if err != nil {
		return err
	}

	return response.Error
}

func (c *client) DesiredLRPs(filter models.DesiredLRPFilter) ([]*models.DesiredLRP, *models.Error) {
	request := models.DesiredLRPsRequest{
		Domain: filter.Domain,
	}
	response := models.DesiredLRPsResponse{}
	err := c.doRequest(DesiredLRPsRoute, nil, nil, &request, &response)
	if err != nil {
		return nil, err
	}

	return response.DesiredLrps, response.Error
}

func (c *client) DesiredLRPByProcessGuid(processGuid string) (*models.DesiredLRP, *models.Error) {
	request := models.DesiredLRPByProcessGuidRequest{
		ProcessGuid: processGuid,
	}
	response := models.DesiredLRPResponse{}
	err := c.doRequest(DesiredLRPByProcessGuidRoute, nil, nil, &request, &response)
	if err != nil {
		return nil, err
	}

	return response.DesiredLrp, response.Error
}

func (c *client) doDesiredLRPLifecycleRequest(route string, request proto.Message) *models.Error {
	response := models.DesiredLRPLifecycleResponse{}
	err := c.doRequest(route, nil, nil, request, &response)
	if err != nil {
		return err
	}
	return response.Error
}

func (c *client) DesireLRP(desiredLRP *models.DesiredLRP) *models.Error {
	request := models.DesireLRPRequest{
		DesiredLrp: desiredLRP,
	}
	return c.doDesiredLRPLifecycleRequest(DesireDesiredLRPRoute, &request)
}

func (c *client) UpdateDesiredLRP(processGuid string, update *models.DesiredLRPUpdate) *models.Error {
	request := models.UpdateDesiredLRPRequest{
		ProcessGuid: processGuid,
		Update:      update,
	}
	return c.doDesiredLRPLifecycleRequest(UpdateDesiredLRPRoute, &request)
}

func (c *client) RemoveDesiredLRP(processGuid string) *models.Error {
	request := models.RemoveDesiredLRPRequest{
		ProcessGuid: processGuid,
	}
	return c.doDesiredLRPLifecycleRequest(RemoveDesiredLRPRoute, &request)
}

func (c *client) ConvergeLRPs() *models.Error {
	route := ConvergeLRPsRoute
	response := models.ConvergeLRPsResponse{}
	err := c.doRequest(route, nil, nil, nil, &response)
	if err != nil {
		return err
	}
	return response.Error
}

func (c *client) Tasks() ([]*models.Task, *models.Error) {
	request := models.TasksRequest{}
	response := models.TasksResponse{}
	err := c.doRequest(TasksRoute, nil, nil, &request, &response)
	if err != nil {
		return nil, err
	}

	return response.Tasks, response.Error
}

func (c *client) TasksByDomain(domain string) ([]*models.Task, *models.Error) {
	request := models.TasksRequest{
		Domain: domain,
	}
	response := models.TasksResponse{}
	err := c.doRequest(TasksRoute, nil, nil, &request, &response)
	if err != nil {
		return nil, err
	}

	return response.Tasks, response.Error
}

func (c *client) TasksByCellID(cellId string) ([]*models.Task, *models.Error) {
	request := models.TasksRequest{
		CellId: cellId,
	}
	response := models.TasksResponse{}
	err := c.doRequest(TasksRoute, nil, nil, &request, &response)
	if err != nil {
		return nil, err
	}

	return response.Tasks, response.Error
}

func (c *client) TaskByGuid(taskGuid string) (*models.Task, *models.Error) {
	request := models.TaskByGuidRequest{
		TaskGuid: taskGuid,
	}
	response := models.TaskResponse{}
	err := c.doRequest(TaskByGuidRoute, nil, nil, &request, &response)
	if err != nil {
		return nil, err
	}

	return response.Task, response.Error
}

func (c *client) doTaskLifecycleRequest(route string, request proto.Message) *models.Error {
	response := models.TaskLifecycleResponse{}
	err := c.doRequest(route, nil, nil, request, &response)
	if err != nil {
		return err
	}
	return response.Error
}

func (c *client) DesireTask(taskGuid, domain string, taskDef *models.TaskDefinition) *models.Error {
	route := DesireTaskRoute
	request := models.DesireTaskRequest{
		TaskGuid:       taskGuid,
		Domain:         domain,
		TaskDefinition: taskDef,
	}
	return c.doTaskLifecycleRequest(route, &request)
}

func (c *client) StartTask(taskGuid string, cellId string) (bool, *models.Error) {
	request := &models.StartTaskRequest{
		TaskGuid: taskGuid,
		CellId:   cellId,
	}
	response := &models.StartTaskResponse{}
	err := c.doRequest(StartTaskRoute, nil, nil, request, response)
	if err != nil {
		return false, err
	}
	return response.ShouldStart, response.Error
}

func (c *client) CancelTask(taskGuid string) *models.Error {
	request := models.TaskGuidRequest{
		TaskGuid: taskGuid,
	}
	route := CancelTaskRoute
	return c.doTaskLifecycleRequest(route, &request)
}

func (c *client) ResolvingTask(taskGuid string) *models.Error {
	request := models.TaskGuidRequest{
		TaskGuid: taskGuid,
	}
	route := ResolvingTaskRoute
	return c.doTaskLifecycleRequest(route, &request)
}

func (c *client) DeleteTask(taskGuid string) *models.Error {
	request := models.TaskGuidRequest{
		TaskGuid: taskGuid,
	}
	route := DeleteTaskRoute
	return c.doTaskLifecycleRequest(route, &request)
}

func (c *client) FailTask(taskGuid, failureReason string) *models.Error {
	request := models.FailTaskRequest{
		TaskGuid:      taskGuid,
		FailureReason: failureReason,
	}
	route := FailTaskRoute
	return c.doTaskLifecycleRequest(route, &request)
}

func (c *client) CompleteTask(taskGuid, cellId string, failed bool, failureReason, result string) *models.Error {
	request := models.CompleteTaskRequest{
		TaskGuid:      taskGuid,
		CellId:        cellId,
		Failed:        failed,
		FailureReason: failureReason,
		Result:        result,
	}
	route := CompleteTaskRoute
	return c.doTaskLifecycleRequest(route, &request)
}

func (c *client) ConvergeTasks(kickTaskDuration, expirePendingTaskDuration, expireCompletedTaskDuration time.Duration) *models.Error {
	request := &models.ConvergeTasksRequest{
		KickTaskDuration:            kickTaskDuration.Nanoseconds(),
		ExpirePendingTaskDuration:   expirePendingTaskDuration.Nanoseconds(),
		ExpireCompletedTaskDuration: expireCompletedTaskDuration.Nanoseconds(),
	}
	response := models.ConvergeTasksResponse{}
	route := ConvergeTasksRoute
	err := c.doRequest(route, nil, nil, request, &response)
	if err != nil {
		return err
	}
	return response.Error
}

func (c *client) SubscribeToEvents() (events.EventSource, *models.Error) {
	eventSource, err := sse.Connect(c.streamingHTTPClient, time.Second, func() *http.Request {
		request, err := c.reqGen.CreateRequest(EventStreamRoute, nil, nil)
		if err != nil {
			panic(err) // totally shouldn't happen
		}

		return request
	})

	if err != nil {
		return nil, models.NewError(models.Error_NetworkError, err.Error())
	}

	return events.NewEventSource(eventSource), nil
}

func (c *client) createRequest(requestName string, params rata.Params, queryParams url.Values, message proto.Message) (*http.Request, *models.Error) {
	var messageBody []byte
	var err error
	if message != nil {
		messageBody, err = proto.Marshal(message)
		if err != nil {
			return nil, models.NewError(models.Error_InvalidProtobufMessage, err.Error())
		}
	}

	request, err := c.reqGen.CreateRequest(requestName, params, bytes.NewReader(messageBody))
	if err != nil {
		return nil, models.NewError(models.Error_InvalidRequest, err.Error())
	}

	request.URL.RawQuery = queryParams.Encode()
	request.ContentLength = int64(len(messageBody))
	request.Header.Set("Content-Type", ProtoContentType)
	return request, nil
}

func (c *client) doEvacRequest(route string, defaultKeepContainer bool, request proto.Message) (bool, *models.Error) {
	var response models.EvacuationResponse
	err := c.doRequest(route, nil, nil, request, &response)
	if err != nil {
		return defaultKeepContainer, err
	}

	return response.KeepContainer, response.Error
}

func (c *client) doRequest(requestName string, params rata.Params, queryParams url.Values, requestBody, responseBody proto.Message) *models.Error {
	request, err := c.createRequest(requestName, params, queryParams, requestBody)
	if err != nil {
		return err
	}
	return c.do(request, responseBody)
}

func (c *client) do(request *http.Request, responseObject proto.Message) *models.Error {
	response, err := c.httpClient.Do(request)
	if err != nil {
		return models.NewError(models.Error_UnknownError, err.Error())
	}
	defer response.Body.Close()

	var parsedContentType string
	if contentType, ok := response.Header[ContentTypeHeader]; ok {
		parsedContentType, _, _ = mime.ParseMediaType(contentType[0])
	}

	if routerError, ok := response.Header[XCfRouterErrorHeader]; ok {
		return models.NewError(models.Error_RouterError, routerError[0])
	}

	if parsedContentType == ProtoContentType {
		return handleProtoResponse(response, responseObject)
	} else {
		return handleNonProtoResponse(response)
	}
}

func handleProtoResponse(response *http.Response, responseObject proto.Message) *models.Error {
	if responseObject == nil {
		return models.NewError(models.Error_InvalidRequest, "responseObject cannot be nil")
	}

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.NewError(models.Error_InvalidResponse, fmt.Sprint("failed to read body: ", err.Error()))
	}

	err = proto.Unmarshal(buf, responseObject)
	if err != nil {
		return models.NewError(models.Error_InvalidProtobufMessage, fmt.Sprintf("failed to unmarshal proto", err.Error()))
	}

	return nil
}

func handleNonProtoResponse(response *http.Response) *models.Error {
	if response.StatusCode > 299 {
		return models.NewError(models.Error_InvalidResponse, fmt.Sprintf("Invalid Response with status code: %d", response.StatusCode))
	}
	return nil
}
