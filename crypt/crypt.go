package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type StreamCrypt struct {
	block cipher.Block
	iv    []byte
}

func NewIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	return iv, nil
}

func NewStreamCrypt(key, iv []byte) (*StreamCrypt, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(iv) != block.BlockSize() {
		return nil, errors.New("iv length does not match block size")
	}

	return &StreamCrypt{
		block: block,
		iv:    iv,
	}, nil
}

func (f *StreamCrypt) Encrypt(ciphertext io.Writer, plaintext io.Reader) error {
	encrypter := cipher.NewCFBEncrypter(f.block, f.iv)

	w := &cipher.StreamWriter{S: encrypter, W: ciphertext}
	_, err := io.Copy(w, plaintext)
	return err
}

func (f *StreamCrypt) Decrypt(plaintext io.Writer, ciphertext io.Reader) error {
	decrypter := cipher.NewCFBDecrypter(f.block, f.iv)

	r := &cipher.StreamReader{S: decrypter, R: ciphertext}
	_, err := io.Copy(plaintext, r)
	return err
}
