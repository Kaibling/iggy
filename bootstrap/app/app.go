package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	apiservice "github.com/kaibling/apiforge/service"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/bootstrap/api"
	bootstrap_broker "github.com/kaibling/iggy/bootstrap/broker"
	"github.com/kaibling/iggy/persistence/psql"
	"github.com/kaibling/iggy/pkg/config"
)

func Run(withWorker bool, withAPI bool) error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	logger := apiservice.BuildLogger(apiservice.LogConfig{
		LogDriver: cfg.App.Logger,
		LogLevel:  "debug",
	})
	logger.AddStringField("scope", "startup")

	ctx, ctxCancel := context.WithCancel(context.Background())

	conn, err := psql.New(ctx, cfg.DB)
	if err != nil {
		ctxCancel()

		return err
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	workerConfig := broker.SubscriberConfig{
		Config:   cfg,
		Username: config.SystemUser,
		DBPool:   conn,
	}

	if withWorker {
		var worker broker.Subscriber

		worker, err = bootstrap_broker.NewSubscriber(workerConfig, "loopback", logger)
		if err != nil {
			logger.Error(err)
		}

		// hopefully not blocking
		go worker.Subscribe(ctx, cfg.Broker.Channel) //nolint: errcheck
		logger.Info("worker started")
	}

	if withAPI {
		if err := api.Start(ctx, cfg, logger, conn); err != nil {
			logger.Error(err)
		}

		logger.Info("api started")
	}

	logger.Info("application started. Ready...")
	<-interrupt

	logger.Info("stopping application")
	ctxCancel()

	gracePeriod := 400 * time.Millisecond //nolint:mnd
	time.Sleep(gracePeriod)

	return nil
}
