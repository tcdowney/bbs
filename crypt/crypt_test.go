package crypt_test

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/cloudfoundry-incubator/bbs/crypt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type C struct {
	A string `json:"a"`
}

type J struct {
	A string `json:"a"`
	B int    `json:"b"`
	C C      `json:"c"`
}

var _ = Describe("StreamCrypt", func() {
	FIt("encrypts and decrypts", func() {
		c, err := crypt.NewStreamCrypt(key, iv)
		Expect(err).NotTo(HaveOccurred())

		// obj := J{A: "asdf", B: 3, C: C{A: "potato"}}
		// jsonValue, err := json.Marshal(&obj)
		jsonValue := []byte("12345678912")
		Expect(err).NotTo(HaveOccurred())

		plaintext := bytes.NewBuffer(jsonValue)
		ciphertext := &bytes.Buffer{}

		encoder := base64.NewEncoder(base64.StdEncoding, ciphertext)
		err = c.Encrypt(encoder, plaintext)
		Expect(err).NotTo(HaveOccurred())

		encrypted := ciphertext.Bytes()
		decrypted := &bytes.Buffer{}

		encBuffer := base64.NewDecoder(base64.StdEncoding, bytes.NewBuffer(encrypted))
		err = c.Decrypt(decrypted, encBuffer)
		Expect(err).NotTo(HaveOccurred())

		fmt.Printf("%s", hex.Dump(encrypted))
		fmt.Printf("%s", hex.Dump(decrypted.Bytes()))

		otherObj := J{}
		err = json.Unmarshal(decrypted.Bytes(), &otherObj)
		Expect(err).NotTo(HaveOccurred())
		// Expect(obj).To(Equal(otherObj))
	})

	Measure("time to encrypt 1k iterations of 1k", func(b Benchmarker) {
		c, err := crypt.NewStreamCrypt(key, iv)
		Expect(err).NotTo(HaveOccurred())

		outputBuffer := make([]byte, len(KILOBYTE_CIPHERTEXT))

		b.Time("encrypt", func() {
			for i := 0; i < 1024; i++ {
				ciphertext := bytes.NewBuffer(outputBuffer)
				err = c.Encrypt(ciphertext, bytes.NewBuffer(KILOBYTE_PLAINTEXT))
				Expect(err).NotTo(HaveOccurred())
			}
		})
	}, 1024)

	Measure("time to decrypt 1k iterations of 1k", func(b Benchmarker) {
		c, err := crypt.NewStreamCrypt(key, iv)
		Expect(err).NotTo(HaveOccurred())

		outputBuffer := make([]byte, len(KILOBYTE_PLAINTEXT))

		b.Time("decrypt", func() {
			for i := 0; i < 1024; i++ {
				plaintext := bytes.NewBuffer(outputBuffer)
				err = c.Decrypt(plaintext, bytes.NewBuffer(KILOBYTE_CIPHERTEXT))
				Expect(err).NotTo(HaveOccurred())
			}
		})
	}, 1024)

	Measure("time to encrypt and encode 1k iterations of 1k", func(b Benchmarker) {
		c, err := crypt.NewStreamCrypt(key, iv)
		Expect(err).NotTo(HaveOccurred())

		outputBuffer := make([]byte, len(KILOBYTE_ENCODED))

		b.Time("encrypt", func() {
			for i := 0; i < 1024; i++ {
				encoded := bytes.NewBuffer(outputBuffer)
				ciphertext := base64.NewEncoder(base64.StdEncoding, encoded)
				err = c.Encrypt(ciphertext, bytes.NewBuffer(KILOBYTE_PLAINTEXT))
				Expect(err).NotTo(HaveOccurred())
			}
		})
	}, 1024)

	Measure("time to decode and decrypt 1k iterations of 1k", func(b Benchmarker) {
		c, err := crypt.NewStreamCrypt(key, iv)
		Expect(err).NotTo(HaveOccurred())

		outputBuffer := make([]byte, len(KILOBYTE_PLAINTEXT))

		b.Time("decode and decrypt", func() {
			for i := 0; i < 1024; i++ {
				encoded := bytes.NewBuffer(KILOBYTE_ENCODED)
				plaintext := bytes.NewBuffer(outputBuffer)
				ciphertext := base64.NewDecoder(base64.StdEncoding, encoded)
				err = c.Decrypt(plaintext, ciphertext)
				Expect(err).NotTo(HaveOccurred())
			}
		})
	}, 1024)

	Measure("time to encrypt 1M", func(b Benchmarker) {
		c, err := crypt.NewStreamCrypt(key, iv)
		Expect(err).NotTo(HaveOccurred())

		outputBuffer := make([]byte, len(MEGABYTE_CIPHERTEXT))

		b.Time("encrypt", func() {
			ciphertext := bytes.NewBuffer(outputBuffer)
			err = c.Encrypt(ciphertext, bytes.NewBuffer(MEGABYTE_PLAINTEXT))
			Expect(err).NotTo(HaveOccurred())
		})
	}, 1024)

	Measure("time to decrypt 1M", func(b Benchmarker) {
		c, err := crypt.NewStreamCrypt(key, iv)
		Expect(err).NotTo(HaveOccurred())

		outputBuffer := make([]byte, len(MEGABYTE_PLAINTEXT))

		b.Time("decrypt", func() {
			plaintext := bytes.NewBuffer(outputBuffer)
			err = c.Decrypt(plaintext, bytes.NewBuffer(MEGABYTE_CIPHERTEXT))
			Expect(err).NotTo(HaveOccurred())
		})
	}, 1)

	Measure("time to encrypt and encode 1M", func(b Benchmarker) {
		c, err := crypt.NewStreamCrypt(key, iv)
		Expect(err).NotTo(HaveOccurred())

		b.Time("encrypt and encode", func() {
			encoded := &bytes.Buffer{}
			ciphertext := base64.NewEncoder(base64.StdEncoding, encoded)
			err = c.Encrypt(ciphertext, bytes.NewBuffer(MEGABYTE_PLAINTEXT))
			Expect(err).NotTo(HaveOccurred())
		})
	}, 1)

	Measure("time to decode and decrypt 1M", func(b Benchmarker) {
		c, err := crypt.NewStreamCrypt(key, iv)
		Expect(err).NotTo(HaveOccurred())

		outputBuffer := make([]byte, len(MEGABYTE_PLAINTEXT))

		b.Time("decode and decrypt", func() {
			encoded := bytes.NewBuffer(MEGABYTE_ENCODED)
			plaintext := bytes.NewBuffer(outputBuffer)
			ciphertext := base64.NewDecoder(base64.StdEncoding, encoded)
			err = c.Decrypt(plaintext, ciphertext)
			Expect(err).NotTo(HaveOccurred())
		})
	}, 1024)
})
