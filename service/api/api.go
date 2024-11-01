package api

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/kaibling/apiforge/apictx"
	"github.com/kaibling/apiforge/handler"
	"github.com/kaibling/apiforge/middleware"
	apiservice "github.com/kaibling/apiforge/service"
	"github.com/kaibling/iggy/api"
	"github.com/kaibling/iggy/config"
	"github.com/kaibling/iggy/initservice"
	"github.com/kaibling/iggy/migration"
	"github.com/kaibling/iggy/persistence/psql"
)

func Start() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	l := apiservice.BuildLogger(cfg.Logger)
	ctx := context.Background()

	ctx = context.WithValue(ctx, apictx.String("logger"), l)
	conn, err := psql.New(cfg)
	if err != nil {
		return err
	}
	ctx = context.WithValue(ctx, apictx.String("db"), conn)
	ctx = context.WithValue(ctx, apictx.String("cfg"), cfg)
	if err = migration.SelfMigrate(cfg); err != nil {
		return err
	} else {
		l.LogLine("rh migration successful")
	}

	userService := initservice.NewUserService(ctx, config.SYSTEM_USER)

	pwd, err := userService.EnsureAdmin(cfg.AdminPassword)
	if err != nil {
		fmt.Println(err.Error())
	}
	if pwd != "" {
		l.LogLine(fmt.Sprintf("Admin Password: %s", pwd))
	}

	root := chi.NewRouter()
	// context
	root.Use(middleware.AddContext("logger", l))
	root.Use(middleware.AddContext("db", apictx.GetValue(ctx, "db")))
	root.Use(middleware.AddContext("cfg", apictx.GetValue(ctx, "cfg")))

	// middleware
	root.Use(middleware.InitEnvelope)
	root.Use(middleware.SaveBody)
	root.Use(middleware.LogRequest)
	root.Use(middleware.Recoverer)

	// mount api endpoint
	root.Mount("/api/v1", api.ApiRoute())

	root.NotFound(handler.NotFound)

	a := apiservice.New(ctx, apiservice.ServerConfig{BindingIP: cfg.BindingIP, BindingPort: cfg.BindingPort})
	a.AddCustomLogger(l)
	a.StartBlocking(root)
	return nil
}
