package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBlockHash(t *testing.T) {
	t.Parallel()

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
		Ts: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix(),
	}

	hash, err := block.Hash()
	assert.Nil(t, err)
	assert.Equal(t, "abf3122c9d666b3e9cd9481c7fe2d78798db7282a4c1d642825c3ff5bd9fcc19", hash)
}

func TestBlockMine(t *testing.T) {
	t.Parallel()

	var (
		err   error
		block = Block{
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
			Ts: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix(),
		}
	)

	block.HashSum, err = block.Mine(2)
	assert.Nil(t, err)
	assert.Equal(t, int64(105), block.Pow)
	assert.Equal(t, "005a96f3290167cd26988b037791fcc249a74b1da028659cf79452bb17c12d44", block.HashSum)
}

func TestValidBlockValidate(t *testing.T) {
	t.Parallel()

	var (
		err   error
		block = Block{
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
			Ts: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix(),
		}
	)

	block.HashSum, err = block.Hash()
	assert.Nil(t, err)

	isValid, err := block.Validate()
	assert.Nil(t, err)
	assert.Equal(t, true, isValid)
}

func TestInvalidBlockValidate(t *testing.T) {
	t.Parallel()

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
		Ts: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix(),
	}

	hash, err := block.Hash()
	assert.Nil(t, err)

	block.HashSum = hash[1:]

	isValid, err := block.Validate()
	assert.Nil(t, err)
	assert.Equal(t, false, isValid)
}
