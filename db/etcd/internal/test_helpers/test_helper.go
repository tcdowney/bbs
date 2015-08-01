package test_helpers

import (
	"github.com/cloudfoundry-incubator/consuladapter"
	etcdclient "github.com/coreos/go-etcd/etcd"
)

func NewTestHelper(etcdClient *etcdclient.Client, consulSession *consuladapter.Session) *TestHelper {
	return &TestHelper{etcdClient: etcdClient, consulSession: consulSession}
}

type TestHelper struct {
	etcdClient    *etcdclient.Client
	consulSession *consuladapter.Session
}
