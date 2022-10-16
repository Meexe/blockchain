package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionHash(t *testing.T) {
	t.Parallel()

	var (
		testTransactionHashResult = []byte{
			0xc, 0x6c, 0xc3, 0xbf, 0x37, 0xf, 0x2a, 0x30, 0xa4, 0x80, 0xfe, 0xdb, 0xe8, 0xf3, 0xef, 0xd5,
			0x85, 0xa4, 0xf1, 0x4e, 0x1f, 0x8d, 0xa3, 0xf6, 0xfa, 0xb7, 0xa9, 0x6d, 0x7a, 0x1, 0x66, 0xa0,
		}
		tx = Transaction{
			[]byte("address"),
			[]byte("address"),
			1.0,
			[]byte("signature"),
		}
	)

	hash, err := tx.Hash()
	assert.Nil(t, err)
	assert.Equal(t, testTransactionHashResult, hash)
}

func TestValidTransactionSign(t *testing.T) {
	t.Parallel()

	key, err := rsa.GenerateKey(rand.Reader, 4096)
	assert.Nil(t, err)

	addr := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	tx := &Transaction{addr, addr, 1.0, nil}

	tx.Signature, err = tx.Sign(key)
	assert.Nil(t, err)

	err = tx.Validate()
	assert.Nil(t, err)
}

func TestInvalidTransactionSign(t *testing.T) {
	t.Parallel()

	key, err := rsa.GenerateKey(rand.Reader, 4096)
	assert.Nil(t, err)

	addr := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	tx := &Transaction{addr, addr, 1.0, nil}

	tx.Signature, err = tx.Sign(key)
	assert.Nil(t, err)
	tx.Signature = tx.Signature[1:]

	err = tx.Validate()
	assert.NotNil(t, err)
}
