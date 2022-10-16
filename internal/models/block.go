package models

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// Block - блок в блокчейне
//easyjson:json
type Block struct {
	HashSum      string       `json:"hash,omitempty"`
	PrevHash     string       `json:"prevHash"`
	Transactions Transactions `json:"transactions"`
	Ts           int64        `json:"ts"`
	Pow          int64        `json:"pow"`
}

func (b Block) Hash() (string, error) {
	b.HashSum = ""
	data, err := b.MarshalJSON()
	if err != nil {
		return "", err
	}

	blockHash := sha256.Sum256(data)
	return fmt.Sprintf("%x", blockHash), nil
}

func (b *Block) Mine(difficulty int) (hash string, err error) {
	for !strings.HasPrefix(string(hash), strings.Repeat("0", difficulty)) {
		b.Pow++
		if hash, err = b.Hash(); err != nil {
			return
		}
	}
	return
}

func (b *Block) Validate() (bool, error) {
	hash, err := b.Hash()
	if err != nil {
		return false, err
	}

	return hash == b.HashSum, nil
}
