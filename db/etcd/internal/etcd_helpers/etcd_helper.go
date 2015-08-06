package etcd_helpers

import (
	"github.com/cloudfoundry-incubator/bbs/crypt"
	etcdclient "github.com/coreos/go-etcd/etcd"

	. "github.com/onsi/gomega"
)

func NewETCDHelper(etcdClient *etcdclient.Client) *ETCDHelper {
	streamCrypt, err := crypt.NewStreamCrypt([]byte("-this is an aes-192 key-"), []byte("abcdabcdabcdabcd"))
	Expect(err).NotTo(HaveOccurred())

	return &ETCDHelper{
		etcdClient:  etcdClient,
		streamCrypt: streamCrypt,
	}
}

type ETCDHelper struct {
	etcdClient  *etcdclient.Client
	streamCrypt *crypt.StreamCrypt
}
