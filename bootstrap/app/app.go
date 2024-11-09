package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	apiservice "github.com/kaibling/apiforge/service"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/adapters/broker/loopback"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/bootstrap/api"
	"github.com/kaibling/iggy/pkg/config"
)

func Run(withWorker bool, withApi bool) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	l := apiservice.BuildLogger(cfg.App.Logger)
	//ctx := context.Background()
	ctx, ctxCancel := context.WithCancel(context.Background())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	if withWorker {
		var worker broker.BrokerAdapter
		if cfg.Broker.Broker == "loopback" {
			// TODO no magic number
			internalChannel := make(chan []byte, 100)
			worker = loopback.NewLoopback(ctx, internalChannel, l)
		} else {
			worker, err = bootstrap.NewWorker(ctx, "loopback")
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		// hopefully not blocking
		go worker.Subscribe(cfg.Broker.Channel)
		l.LogLine("worker started")
	}
	if withApi {
		if err := api.Start(cfg, ctx, l); err != nil {
			fmt.Println(err)
		}
		l.LogLine("api started")
	}
	l.LogLine("application started. Ready...")

	<-interrupt
	l.LogLine("stopping application")
	ctxCancel()
	// TODO remove
	time.Sleep(400 * time.Millisecond)
	return nil
}
