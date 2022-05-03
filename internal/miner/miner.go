package Miner

import (
	"context"
)

type miner struct{}

func (m *miner) Start(ctx context.Context) chan error {
	var ch = make(chan error)

	go func(ch chan error) {
		for {
			select {
			case ch <- mine(ctx):
			case <-ctx.Done():
				return
			}
		}
	}(ch)

	return ch
}

func mine(ctx context.Context) error {
	return nil
}
