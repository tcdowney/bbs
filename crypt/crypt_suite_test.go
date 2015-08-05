package crypt_test

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"io"

	"github.com/cloudfoundry-incubator/bbs/crypt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCrypt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crypt Suite")
}

var KILOBYTE_PLAINTEXT []byte
var KILOBYTE_CIPHERTEXT []byte
var KILOBYTE_ENCODED []byte

var MEGABYTE_PLAINTEXT []byte
var MEGABYTE_CIPHERTEXT []byte
var MEGABYTE_ENCODED []byte

var iv []byte
var key []byte

var _ = BeforeSuite(func() {
	var err error

	iv, err = crypt.NewIV()
	Expect(err).NotTo(HaveOccurred())

	key = []byte("-this is an aes-192 key-")
	c, err := crypt.NewStreamCrypt(key, iv)
	Expect(err).NotTo(HaveOccurred())

	KILOBYTE_PLAINTEXT = make([]byte, 1024)
	_, err = io.ReadFull(rand.Reader, KILOBYTE_PLAINTEXT)
	Expect(err).NotTo(HaveOccurred())

	ciphertext := &bytes.Buffer{}
	err = c.Encrypt(ciphertext, bytes.NewBuffer(KILOBYTE_PLAINTEXT))
	Expect(err).NotTo(HaveOccurred())
	KILOBYTE_CIPHERTEXT = ciphertext.Bytes()

	encoded := &bytes.Buffer{}
	encodeWriter := base64.NewEncoder(base64.StdEncoding, encoded)
	_, err = io.Copy(encodeWriter, bytes.NewReader(KILOBYTE_CIPHERTEXT))
	KILOBYTE_ENCODED = encoded.Bytes()

	MEGABYTE_PLAINTEXT = make([]byte, 1024*1024)
	_, err = io.ReadFull(rand.Reader, MEGABYTE_PLAINTEXT)
	Expect(err).NotTo(HaveOccurred())

	ciphertext = &bytes.Buffer{}
	err = c.Encrypt(ciphertext, bytes.NewBuffer(MEGABYTE_PLAINTEXT))
	Expect(err).NotTo(HaveOccurred())
	MEGABYTE_CIPHERTEXT = ciphertext.Bytes()

	encoded = &bytes.Buffer{}
	encodeWriter = base64.NewEncoder(base64.StdEncoding, encoded)
	_, err = io.Copy(encodeWriter, bytes.NewReader(MEGABYTE_CIPHERTEXT))
	MEGABYTE_ENCODED = encoded.Bytes()
})
