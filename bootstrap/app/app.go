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

const innerChannelSize = 100

func Run(withWorker bool, withAPI bool) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	logger := apiservice.BuildLogger(cfg.App.Logger)
	// ctx := context.Background()
	ctx, ctxCancel := context.WithCancel(context.Background())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	if withWorker {
		var worker broker.BrokerAdapter

		if cfg.Broker.Broker == "loopback" {

			internalChannel := make(chan []byte, innerChannelSize)
			worker = loopback.NewLoopback(ctx, internalChannel, logger)
		} else {
			worker, err = bootstrap.NewWorker(ctx, "loopback")
			if err != nil {
				fmt.Println(err.Error())
			}
		}

		// hopefully not blocking
		go worker.Subscribe(cfg.Broker.Channel)
		logger.LogLine("worker started")
	}

	if withAPI {
		if err := api.Start(ctx, cfg, logger); err != nil {
			fmt.Println(err)
		}

		logger.LogLine("api started")
	}

	logger.LogLine("application started. Ready...")

	<-interrupt

	logger.LogLine("stopping application")

	ctxCancel()
	// TODO context should be with timeout
	gracePeriod := 400 * time.Millisecond //nolint:gomnd,mnd
	time.Sleep(gracePeriod)

	return nil
}
