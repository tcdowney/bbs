package db

import (
	"encoding/json"
	"path"
	"time"

	"github.com/cloudfoundry-incubator/bbs"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/pivotal-golang/lager"
)

const maxDesiredLRPGetterWorkPoolSize = 50
const DesiredLRPSchemaRoot = DataSchemaRoot + "desired"

func DesiredLRPSchemaPath(lrp models.DesiredLRP) string {
	return DesiredLRPSchemaPathByProcessGuid(lrp.GetProcessGuid())
}

func DesiredLRPSchemaPathByProcessGuid(processGuid string) string {
	return path.Join(DesiredLRPSchemaRoot, processGuid)
}

func (db *ETCDDB) DesiredLRPs(filter models.DesiredLRPFilter, logger lager.Logger) (*models.DesiredLRPs, *bbs.Error) {
	start := time.Now()
	root, bbsErr := db.fetchRecursiveRaw(DesiredLRPSchemaRoot, logger)
	end := time.Now()
	if bbsErr.Equal(bbs.ErrResourceNotFound) {
		return &models.DesiredLRPs{}, nil
	}
	if bbsErr != nil {
		return nil, bbsErr
	}
	if root.Nodes.Len() == 0 {
		return &models.DesiredLRPs{}, nil
	}

	logger.Info("etc fetch recursive: " + end.Sub(start).String())
	desiredLRPs := models.DesiredLRPs{}

	var base64time int64

	logger.Debug("performing-deserialization-work")
	start = time.Now()

	for _, node := range root.Nodes {
		node := node

		var lrp models.DesiredLRP
		logger.Debug("logging node value", lager.Data{"node value": node.Value})
		// start := time.Now()
		// m, err := base64.StdEncoding.DecodeString(node.Value)
		// end := time.Now()
		// if err != nil {
		// 	logger.Error("failed-parsing-desired-lrp", err)
		// 	return nil, bbs.ErrUnknownError
		// }

		// err = lrp.Unmarshal([]byte(m))
		//err := models.FromJSON([]byte(node.Value), &lrp)
		err := json.Unmarshal([]byte(node.Value), &lrp)
		if err != nil {
			logger.Error("failed-parsing-desired-lrp", err)
			return nil, bbs.ErrUnknownError
		}

		if filter.Domain == "" || lrp.GetDomain() == filter.Domain {

			desiredLRPs.DesiredLrps = append(desiredLRPs.DesiredLrps, &lrp)
			// base64time += int64(end.Sub(start))
		}
	}

	end = time.Now()

	logger.Info("decode+unmarshal", lager.Data{"total": end.Sub(start).String(), "base64": base64time})

	logger.Debug("succeeded-performing-deserialization-work", lager.Data{"num-desired-lrps": len(desiredLRPs.GetDesiredLrps())})

	return &desiredLRPs, nil
}

func (db *ETCDDB) DesiredLRPByProcessGuid(processGuid string, logger lager.Logger) (*models.DesiredLRP, *bbs.Error) {
	node, bbsErr := db.fetchRaw(DesiredLRPSchemaPathByProcessGuid(processGuid), logger)
	if bbsErr != nil {
		return nil, bbsErr
	}

	var lrp models.DesiredLRP
	deserializeErr := models.FromJSON([]byte(node.Value), &lrp)
	if deserializeErr != nil {
		logger.Error("failed-parsing-desired-lrp", deserializeErr)
		return nil, bbs.ErrDeserializeJSON
	}

	return &lrp, nil
}
