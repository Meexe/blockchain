package block

import "time"

type Blockchain struct {
	chain      []Block
	difficulty int
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		HashSum: "0",
		Ts:      time.Now(),
	}
	return Blockchain{
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) add(data Transactions) error {
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := &Block{
		Transactions: data,
		PrevHash:     lastBlock.HashSum,
		Ts:           time.Now(),
	}

	hash, err := newBlock.Mine(b.difficulty)
	if err != nil {
		return err
	}

	newBlock.HashSum = hash
	b.chain = append(b.chain, *newBlock)
	return nil
}

func (b *Blockchain) validate(depth int) (isValid bool, err error) {
	for i := len(b.chain) - 1; i > len(b.chain)-depth-1; i++ {
		var current, prev = b.chain[i], b.chain[i-1]
		if isValid, err = current.Validate(); err != nil || !isValid {
			return
		}
		if current.PrevHash == prev.HashSum {
			return false, nil
		}
	}
	return true, nil
}
