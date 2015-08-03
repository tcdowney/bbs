package etcd

import (
	"sync"

	"github.com/cloudfoundry-incubator/auctioneer"

	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/coreos/go-etcd/etcd"
	"github.com/pivotal-golang/clock"
	"github.com/pivotal-golang/lager"
)

const DataSchemaRoot = "/v1/"

const (
	ETCDErrKeyNotFound  = 100
	ETCDErrKeyExists    = 105
	ETCDErrIndexCleared = 401
)

type ETCDDB struct {
	client            *etcd.Client
	clock             clock.Clock
	inflightWatches   map[chan bool]bool
	inflightWatchLock *sync.Mutex
	auctioneerClient  auctioneer.Client
}

func NewETCD(etcdClient *etcd.Client, auctioneerClient auctioneer.Client, clock clock.Clock) *ETCDDB {
	return &ETCDDB{etcdClient,
		clock,
		map[chan bool]bool{},
		&sync.Mutex{},
		auctioneerClient,
	}
}

func (db *ETCDDB) fetchRecursiveRaw(logger lager.Logger, key string) (*etcd.Node, *models.Error) {
	logger.Debug("fetching-recursive-from-etcd")
	response, err := db.client.Get(key, false, true)
	if err != nil {
		return nil, ErrorFromEtcdError(logger, err)
	}
	logger.Debug("succeeded-fetching-recursive-from-etcd", lager.Data{"num-lrps": response.Node.Nodes.Len()})
	return response.Node, nil
}

func (db *ETCDDB) fetchRaw(logger lager.Logger, key string) (*etcd.Node, *models.Error) {
	logger.Debug("fetching-from-etcd")
	response, err := db.client.Get(key, false, false)
	if err != nil {
		return nil, ErrorFromEtcdError(logger, err)
	}
	logger.Debug("succeeded-fetching-from-etcd")
	return response.Node, nil
}

func ErrorFromEtcdError(logger lager.Logger, err error) *models.Error {
	switch etcdErrCode(err) {
	case ETCDErrKeyNotFound:
		logger.Debug("no-node-to-fetch")
		return models.ErrResourceNotFound
	case ETCDErrKeyExists:
		return models.ErrResourceExists
	default:
		logger.Error("failed-fetching-from-etcd", err)
		return models.ErrUnknownError
	}
}

func etcdErrCode(err error) int {
	if err != nil {
		switch err.(type) {
		case etcd.EtcdError:
			return err.(etcd.EtcdError).ErrorCode
		case *etcd.EtcdError:
			return err.(*etcd.EtcdError).ErrorCode
		}
	}
	return 0
}
