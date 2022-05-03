package listener

import "context"

type listener struct{}

func New() *listener {
	return &listener{}
}

func (l *listener) Start(ctx context.Context) chan error {
	return make(chan error)
}
