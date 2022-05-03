package node

import (
	"context"

	"github.com/Meexe/blockchain/tools"
)

type logger interface {
	Info()
	Warn()
	Error()
	Fatal()
	Stop()
}

type miner interface {
	Start(context.Context) chan error
	Stop()
}

type listener interface {
	Start(context.Context) chan error
	Stop()
}

type sender interface {
	Send(context.Context)
}

type node struct {
	miner    miner
	listener listener
	sender   sender
	logger   logger
	cancel   context.CancelFunc
	ch       chan error
}

func New() *node {
	return &node{}
}

func (n *node) Start(ctx context.Context) {
	ctx, n.cancel = context.WithCancel(ctx)
	n.ch = tools.Merge(
		n.miner.Start(ctx),
		n.listener.Start(ctx),
	)

}

func (n node) Stop() {
	close(n.ch)
	n.cancel()
}
