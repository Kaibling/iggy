package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	apiservice "github.com/kaibling/apiforge/service"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/bootstrap/api"
	bootstrap_broker "github.com/kaibling/iggy/bootstrap/broker"
	"github.com/kaibling/iggy/persistence/psql"
	broker_server "github.com/kaibling/iggy/pkg/broker"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/service"
)

const shutdownTime = 100 * time.Millisecond

func Run(withWorker bool, withAPI bool, version, buildTime string) error { //nolint: funlen
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	logger := apiservice.BuildLogger(apiservice.LogConfig{
		LogDriver:    cfg.App.Logger,
		LogLevel:     "debug",
		RequestBody:  false,
		ResponseBody: false,
		JSON:         false,
	})
	logger.AddStringField("scope", "startup")
	logger.Info("version: " + version)
	logger.Info("build time: " + buildTime)

	ctx, ctxCancel := context.WithCancel(context.Background())

	conn, err := psql.New(ctx, cfg.DB)
	if err != nil {
		ctxCancel()

		return err
	}

	sConfig := service.Config{
		Config:   cfg,
		Username: config.SystemUser,
		DBPool:   conn,
		Log:      logger,
	}

	// load workflows
	wfs, err := bootstrap.NewWorkflowService(ctx, sConfig, "export")
	if err != nil {
		ctxCancel()

		return err
	}

	if err := wfs.ImportFromFiles(cfg.App.ExportLocalPath); err != nil {
		ctxCancel()

		return err
	}

	// if err := wfs.ExportToGit(cfg.App.ExportLocalPath, cfg.App.GitToken); err != nil {
	// 	ctxCancel()

	// 	return err
	// }

	// if err := wfs.LoadFromGit(cfg.App.ImportLocalPath); err != nil {
	// 	if err := git.Clone(cfg.App.ImportRepoURL, cfg.App.ImportLocalPath); err != nil {
	// 		ctxCancel()

	// 		return err
	// 	}
	// }

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	if withWorker {
		// start broker server
		broker_server.StartServer(logger, cfg)
		time.Sleep(shutdownTime)

		// start broker client
		var worker broker.Subscriber

		worker, err = bootstrap_broker.NewSubscriber(sConfig, logger)
		if err != nil {
			logger.Error(err)
		}

		go func() {
			// todo in a loop to retry
			if err := worker.Subscribe(ctx, cfg.Broker.Channel); err != nil {
				logger.Error(err)
			} else {
				logger.Info("worker started")
			}
		}()
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
