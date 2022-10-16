package models

import "time"

// Blockchain - блокчейн
type Blockchain struct {
	Chain      []Block `json:"chain"`
	Difficulty int     `json:"diff"`
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		HashSum: "0",
		Ts:      time.Now().Unix(),
	}
	return Blockchain{
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) add(data Transactions) error {
	lastBlock := b.Chain[len(b.Chain)-1]
	newBlock := &Block{
		Transactions: data,
		PrevHash:     lastBlock.HashSum,
		Ts:           time.Now().Unix(),
	}

	hash, err := newBlock.Mine(b.Difficulty)
	if err != nil {
		return err
	}

	newBlock.HashSum = hash
	b.Chain = append(b.Chain, *newBlock)
	return nil
}

func (b *Blockchain) validate(depth int) (isValid bool, err error) {
	for i := len(b.Chain) - 1; i > len(b.Chain)-depth-1; i++ {
		var current, prev = b.Chain[i], b.Chain[i-1]
		if isValid, err = current.Validate(); err != nil || !isValid {
			return
		}
		if current.PrevHash == prev.HashSum {
			return false, nil
		}
	}
	return true, nil
}
