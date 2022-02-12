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
	servers []server
}

func NewApp(servers ...server)  *app {
	return &app{
		servers: servers,
	}
}
func (a *app) Run() error {

	appCtx, cancel:= context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(appCtx)

	for _, srv := range a.servers {
		srv := srv	//这里要格外注意，不解释了，找bug找了半天 MD。。。
		g.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		g.Go(func() error {
			return srv.Start()
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
				cancel()
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

