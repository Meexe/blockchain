package block

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var block = Block{
	PrevHash: "previous hash",
	Transactions: Transactions{
		{
			From:   []byte("vaca"),
			To:     []byte("baca"),
			Amount: 1.0,
		},
		{
			From:   []byte("vaca"),
			To:     []byte("baca"),
			Amount: 10.0,
		},
		{
			From:   []byte("baca"),
			To:     []byte("vaca"),
			Amount: 5.0,
		},
	},
	Ts: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
}

func TestBlockHash(t *testing.T) {
	hash, err := block.Hash()
	assert.Nil(t, err)
	assert.Equal(t, "3bc00d23d161001905a4b9ffb2565b0321e98f9a7a1485137dd38812c4819e4b", hash)
}

func TestBlockMine(t *testing.T) {
	var err error
	block.HashSum, err = block.Mine(1)
	assert.Nil(t, err)
	assert.Equal(t, 25, block.Pow)
	assert.Equal(t, "0e4866e408be65ed2b610c45007854c098425142a551c2c000106e9df32b7649", block.HashSum)
}

func TestValidBlockValidate(t *testing.T) {
	var err error
	block.HashSum, err = block.Hash()
	assert.Nil(t, err)

	isValid, err := block.Validate()
	assert.Nil(t, err)
	assert.Equal(t, true, isValid)
}

func TestInvalidBlockValidate(t *testing.T) {
	hash, err := block.Hash()
	assert.Nil(t, err)

	block.HashSum = hash[1:]

	isValid, err := block.Validate()
	assert.Nil(t, err)
	assert.Equal(t, false, isValid)
}
