package api

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/handler"
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/apiforge/middleware"
	apiservice "github.com/kaibling/apiforge/service"
	"github.com/kaibling/apiforge/status"
	"github.com/kaibling/iggy/api"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/migration"
	"github.com/kaibling/iggy/persistence/psql"
	"github.com/kaibling/iggy/pkg/config"
)

func Start(ctx context.Context, cfg config.Configuration, logger logging.LogWriter) error {
	ctx = context.WithValue(ctx, ctxkeys.LoggerKey, logger)
	conn, err := psql.New(ctx, cfg.DB)
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, ctxkeys.DBConnKey, conn)
	ctx = context.WithValue(ctx, ctxkeys.String("cfg"), cfg)
	if err = migration.SelfMigrate(cfg.DB); err != nil {
		return err
	}

	logger.LogLine("rh migration successful")

	userService := bootstrap.NewUserServiceAnonym(ctx, config.SystemUser)

	pwd, err := userService.EnsureAdmin(cfg.App.AdminPassword)
	if err != nil {
		fmt.Println(err.Error())
	}

	if pwd != "" {
		logger.LogLine("Admin Password: " + pwd)
	}

	root := chi.NewRouter()
	// context
	root.Use(middleware.AddContext("logger", logger))
	root.Use(middleware.AddContext("db", conn))
	root.Use(middleware.AddContext("cfg", cfg))

	// middleware
	root.Use(middleware.InitEnvelope)
	root.Use(middleware.SaveBody)
	root.Use(middleware.LogRequest)
	root.Use(middleware.Recoverer)

	// mount api endpoint
	root.Mount("/api/v1", api.Route())

	root.NotFound(handler.NotFound)

	apiServer := apiservice.New(ctx, apiservice.ServerConfig{
		BindingIP:   cfg.App.BindingIP,
		BindingPort: cfg.App.BindingPort,
	})

	apiServer.AddCustomLogger(logger)
	status.IsReady.Store(true)

	return apiServer.Start(root)
}
