package storage

import (
	"testing"
	"time"

	"github.com/Meexe/blockchain/internal/models"
	"github.com/stretchr/testify/require"
)

func TestSaveLoadBlock(t *testing.T) {

	var (
		err      error
		outBlock models.Block
		inBlock  = models.Block{
			PrevHash: "previous hash",
			Transactions: models.Transactions{
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

	s := New("./testdata/")

	err = s.SaveBlock(inBlock)
	require.Nil(t, err)

	outBlock, err = s.LoadBlock(inBlock.Ts)
	require.Nil(t, err)
	require.Equal(t, inBlock, outBlock)
}

func TestSaveLoadBlockchain(t *testing.T) {

	var (
		err      error
		outChain models.Blockchain
		inChain  = models.Blockchain{
			Chain: []models.Block{
				{
					PrevHash: "genesis hash",
					Transactions: models.Transactions{
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
				},
				{
					PrevHash: "hash 1",
					Transactions: models.Transactions{
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
					Ts: time.Date(2000, time.January, 1, 0, 0, 1, 0, time.UTC).Unix(),
				},
				{
					PrevHash: "hash 2",
					Transactions: models.Transactions{
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
					Ts: time.Date(2000, time.January, 1, 0, 0, 2, 0, time.UTC).Unix(),
				},
			},
			Difficulty: 0,
		}
	)

	s := New("./testdata/")

	err = s.SaveBlockchain(inChain)
	require.Nil(t, err)

	outChain, err = s.LoadBlockchain()
	require.Nil(t, err)
	require.Equal(t, inChain, outChain)
}
