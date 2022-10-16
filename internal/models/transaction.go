package models

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
)

//easyjson:json
type Transactions []Transaction

//easyjson:json
type Transaction struct {
	From      []byte  `json:"from"`
	To        []byte  `json:"to"`
	Amount    float64 `json:"amount"`
	Signature []byte  `json:"sign,omitempty"`
}

func (t Transaction) Hash() ([]byte, error) {
	t.Signature = nil
	body, err := t.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var hash = sha256.New()
	if _, err = hash.Write(body); err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}

func (t *Transaction) Sign(key *rsa.PrivateKey) ([]byte, error) {
	hash, err := t.Hash()
	if err != nil {
		return nil, err
	}

	return rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hash)
}

func (t *Transaction) Validate() error {
	hash, err := t.Hash()
	if err != nil {
		return err
	}

	key, err := x509.ParsePKCS1PublicKey(t.From)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(key, crypto.SHA256, hash, t.Signature)
}
