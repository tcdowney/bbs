package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/gogo/protobuf/proto"
	"github.com/pivotal-golang/lager"
)

type EvacuationHandler struct {
	db     db.EvacuationDB
	logger lager.Logger
}

func NewEvacuationHandler(logger lager.Logger, db db.EvacuationDB) *EvacuationHandler {
	return &EvacuationHandler{
		db:     db,
		logger: logger.Session("evacuation-handler"),
	}
}

type MessageValidator interface {
	proto.Message
	Validate() error
}

func parseRequest(logger lager.Logger, w http.ResponseWriter, req *http.Request, request MessageValidator) bool {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("failed-to-read-body", err)
		writeInternalServerErrorResponse(w, err)
		return false
	}

	err = request.Unmarshal(data)
	if err != nil {
		logger.Error("failed-to-parse-request-body", err)
		writeBadRequestResponse(w, models.InvalidRequest, err)
		return false
	}

	logger.Debug("parsed-request-body", lager.Data{"request": request})
	if err := request.Validate(); err != nil {
		logger.Error("invalid-request", err)
		writeBadRequestResponse(w, models.InvalidRequest, err)
		return false
	}
	return true
}

func (h *EvacuationHandler) RemoveEvacuatingActualLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("remove-evacuating-actual-lrp")

	request := &models.RemoveEvacuatingActualLRPRequest{}
	if !parseRequest(logger, w, req, request) {
		return
	}

	bbsErr := h.db.RemoveEvacuatingActualLRP(logger, request)
	if bbsErr != nil {
		logger.Error("failed-to-remove-evacuating-actual-lrp", bbsErr)
		if bbsErr.Equal(models.ErrResourceNotFound) {
			writeNotFoundResponse(w, bbsErr)
		} else {
			writeInternalServerErrorResponse(w, bbsErr)
		}
		return
	}

	writeEmptyResponse(w, http.StatusNoContent)
}

func (h *EvacuationHandler) EvacuateClaimedActualLRP(w http.ResponseWriter, req *http.Request) {
	logger := h.logger.Session("evacuate-claimed-actual-lrp")

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("failed-to-read-body", err)
		writeInternalServerErrorResponse(w, err)
		return
	}

	request := &models.EvacuateClaimedActualLRPRequest{}
	err = request.Unmarshal(data)
	if err != nil {
		logger.Error("failed-to-parse-request-body", err)
		writeBadRequestResponse(w, models.InvalidRequest, err)
		return
	}

	logger.Debug("parsed-request-body", lager.Data{"request": request})
	if err := request.Validate(); err != nil {
		logger.Error("invalid-request", err)
		writeBadRequestResponse(w, models.InvalidRequest, err)
		return
	}

	bbsErr := h.db.RemoveEvacuatingActualLRP(logger, request)
	if bbsErr != nil {
		logger.Error("failed-to-retire-actual-lrp", bbsErr)
		if bbsErr.Equal(models.ErrResourceNotFound) {
			writeNotFoundResponse(w, bbsErr)
		} else {
			writeInternalServerErrorResponse(w, bbsErr)
		}
		return
	}

	writeEmptyResponse(w, http.StatusNoContent)
}
