package etcd

import (
	"bytes"
	"encoding/base64"
	"sync"

	"github.com/cloudfoundry-incubator/bbs/auctionhandlers"
	"github.com/cloudfoundry-incubator/bbs/cellhandlers"
	"github.com/cloudfoundry-incubator/bbs/crypt"
	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/coreos/go-etcd/etcd"
	"github.com/pivotal-golang/clock"
	"github.com/pivotal-golang/lager"
)

const DataSchemaRoot = "/v1/"

const (
	ETCDErrKeyNotFound  = 100
	ETCDErrIndexCleared = 401
)

type ETCDDB struct {
	client            *etcd.Client
	clock             clock.Clock
	inflightWatches   map[chan bool]bool
	inflightWatchLock *sync.Mutex
	auctioneerClient  auctionhandlers.Client
	cellClient        cellhandlers.Client

	cellDB db.CellDB

	streamCrypt *crypt.StreamCrypt
}

func NewETCD(etcdClient *etcd.Client, auctioneerClient auctionhandlers.Client, cellClient cellhandlers.Client, cellDB db.CellDB, clock clock.Clock) *ETCDDB {
	iv := []byte("abcdabcdabcdabcd")
	streamCrypt, err := crypt.NewStreamCrypt([]byte("-this is an aes-192 key-"), iv)
	if err != nil {
		return nil
	}
	return &ETCDDB{etcdClient,
		clock,
		map[chan bool]bool{},
		&sync.Mutex{},
		auctioneerClient,
		cellClient,
		cellDB,
		streamCrypt,
	}
}

func (db *ETCDDB) fetchRecursiveRaw(logger lager.Logger, key string) (*etcd.Node, *models.Error) {
	logger.Debug("fetching-recursive-from-etcd")
	response, err := db.client.Get(key, false, true)
	if etcdErrCode(err) == ETCDErrKeyNotFound {
		logger.Debug("no-nodes-to-fetch")
		return nil, models.ErrResourceNotFound
	} else if err != nil {
		logger.Error("failed-fetching-recursive-from-etcd", err)
		return nil, models.ErrUnknownError
	}
	logger.Debug("succeeded-fetching-recursive-from-etcd", lager.Data{"num-lrps": response.Node.Nodes.Len()})
	return response.Node, nil
}

func (db *ETCDDB) fetchRaw(logger lager.Logger, key string) (*etcd.Node, *models.Error) {
	logger.Debug("fetching-from-etcd")
	response, err := db.client.Get(key, false, false)
	if etcdErrCode(err) == ETCDErrKeyNotFound {
		logger.Debug("no-node-to-fetch")
		return nil, models.ErrResourceNotFound
	} else if err != nil {
		logger.Error("failed-fetching-from-etcd", err)
		return nil, models.ErrUnknownError
	}
	logger.Debug("succeeded-fetching-from-etcd")
	return response.Node, nil
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

func (db *ETCDDB) decodeFromStorage(s string) []byte {
	if db.streamCrypt != nil {
		fromStorage := bytes.NewBufferString(s)
		toRead := &bytes.Buffer{}
		encrypted := base64.NewDecoder(base64.StdEncoding, fromStorage)
		err := db.streamCrypt.Decrypt(toRead, encrypted)
		if err != nil {
			panic("failed to decrypt")
		}
		return toRead.Bytes()
	} else {
		return []byte(s)
	}
}

func (db *ETCDDB) encodeForStorage(s []byte) string {
	if db.streamCrypt != nil {
		value := bytes.NewBuffer(s)
		toWrite := &bytes.Buffer{}
		encoder := base64.NewEncoder(base64.StdEncoding, toWrite)
		err := db.streamCrypt.Encrypt(encoder, value)
		if err != nil {
			panic("failed to encrypt")
		}
		return toWrite.String()
	} else {
		return string(s)
	}
}
