package app

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type server interface {
	Start() error
	Stop() error
}

type app struct {
	ctx    context.Context
	cancel func()

	servers []server
}

func (a *app) AddServer(servers ...server) {
	a.servers = servers
}
func (a *app) Run() error {

	g, ctx := errgroup.WithContext(a.ctx)

	for _, srv := range a.servers {
		srv := srv	//这里要格外注意，不解释了，找bug找了半天 MD。。。
		g.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		g.Go(func() error {
			err := srv.Start()
			return err
		})
	}

	//信号监控处理
	s := make(chan os.Signal)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-s:
				a.Stop()
			}
		}
	})
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			return err
		}
	}
	return nil

}
func (a *app) Stop() {
	a.cancel()
}

func NewApp() *app {
	var app app
	app.ctx, app.cancel = context.WithCancel(context.Background())
	return &app
}
