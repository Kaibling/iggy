package api

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/handler"
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/apiforge/middleware"
	apiservice "github.com/kaibling/apiforge/service"
	"github.com/kaibling/apiforge/status"
	"github.com/kaibling/iggy/api"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/migration"
	"github.com/kaibling/iggy/pkg/config"
)

func Start(ctx context.Context, cfg config.Configuration, logger logging.Writer, conn *pgxpool.Pool) error {
	logger = logger.NewScope("api_startup")

	ctx = context.WithValue(ctx, ctxkeys.LoggerKey, logger)
	ctx = context.WithValue(ctx, ctxkeys.DBConnKey, conn)
	ctx = context.WithValue(ctx, ctxkeys.AppConfigKey, cfg)

	if err := migration.SelfMigrate(cfg.DB); err != nil {
		return err
	}

	logger.Info("rh migration successful")

	userService, err := bootstrap.NewUserServiceAnonym(ctx, config.SystemUser)
	if err != nil {
		logger.Error(err)

		return err
	}

	pwd, err := userService.EnsureAdmin(cfg.App.AdminPassword)
	if err != nil {
		logger.Error(err)

		return err
	}

	if pwd != "" {
		logger.Info("Admin Password: " + pwd)
	}

	root := chi.NewRouter()
	// context
	root.Use(middleware.AddContext(ctxkeys.LoggerKey, logger))
	root.Use(middleware.AddContext(ctxkeys.DBConnKey, conn))
	root.Use(middleware.AddContext(ctxkeys.AppConfigKey, cfg))

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
		LogLevel:    "debug",
	})

	apiServer.AddCustomLogger(logger)
	status.IsReady.Store(true)

	return apiServer.Start(root)
}
