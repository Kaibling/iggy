package bootstrap

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/apperror"
	repo "github.com/kaibling/iggy/persistence/repository"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/service"
)

func contextDefaultData(ctx context.Context) (config.Configuration, *pgxpool.Pool, string, error) {
	cfg, dbPool, err := contextBasisData(ctx)
	if err != nil {
		return cfg, dbPool, "", apperror.ErrMissingContext
	}

	username, ok := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	if !ok {
		return cfg, dbPool, "", apperror.ErrMissingContext
	}

	return cfg, dbPool, username, nil
}

func contextBasisData(ctx context.Context) (config.Configuration, *pgxpool.Pool, error) {
	cfg, ok := ctxkeys.GetValue(ctx, ctxkeys.AppConfigKey).(config.Configuration)
	if !ok {
		return config.Configuration{}, nil, apperror.ErrMissingContext
	}

	dbPool, ok := ctxkeys.GetValue(ctx, ctxkeys.DBConnKey).(*pgxpool.Pool)
	if !ok {
		return config.Configuration{}, nil, apperror.ErrMissingContext
	}

	return cfg, dbPool, nil
}

func NewUserService(ctx context.Context) (*service.UserService, error) {
	cfg, dbPool, username, err := contextDefaultData(ctx)
	if err != nil {
		return nil, apperror.ErrMissingContext
	}

	userRepo := repo.NewUserRepo(ctx, username, dbPool)

	return service.NewUserService(ctx, userRepo, cfg), nil
}

func NewTokenService(ctx context.Context) (*service.TokenService, error) {
	cfg, dbPool, username, err := contextDefaultData(ctx)
	if err != nil {
		return nil, apperror.ErrMissingContext
	}

	tokenRepo := repo.NewTokenRepo(ctx, username, dbPool)

	return service.NewTokenService(ctx, tokenRepo, cfg), nil
}

func NewUserServiceAnonym(ctx context.Context, username string) (*service.UserService, error) {
	cfg, ok := ctxkeys.GetValue(ctx, ctxkeys.AppConfigKey).(config.Configuration)
	if !ok {
		return nil, apperror.ErrMissingContext
	}

	dbPool, ok := ctxkeys.GetValue(ctx, ctxkeys.DBConnKey).(*pgxpool.Pool)
	if !ok {
		return nil, apperror.ErrMissingContext
	}

	userRepo := repo.NewUserRepo(ctx, username, dbPool)

	return service.NewUserService(ctx, userRepo, cfg), nil
}

func NewTokenServiceAnonym(ctx context.Context, username string) (*service.TokenService, error) {
	cfg, ok := ctxkeys.GetValue(ctx, ctxkeys.AppConfigKey).(config.Configuration)
	if !ok {
		return nil, apperror.ErrMissingContext
	}

	dbPool, ok := ctxkeys.GetValue(ctx, ctxkeys.DBConnKey).(*pgxpool.Pool)
	if !ok {
		return nil, apperror.ErrMissingContext
	}

	tokenRepo := repo.NewTokenRepo(ctx, username, dbPool)

	return service.NewTokenService(ctx, tokenRepo, cfg), nil
}

func NewWorkflowService(ctx context.Context) (*service.WorkflowService, error) {
	cfg, dbPool, username, err := contextDefaultData(ctx)
	if err != nil {
		return nil, apperror.ErrMissingContext
	}

	workflowRepo := repo.NewWorkflowRepo(ctx, username, dbPool)

	return service.NewWorkflowService(ctx, workflowRepo, cfg), nil
}

func NewRunService(ctx context.Context) (*service.RunService, error) {
	cfg, dbPool, username, err := contextDefaultData(ctx)
	if err != nil {
		return nil, apperror.ErrMissingContext
	}

	requestID, ok := ctxkeys.GetValue(ctx, ctxkeys.RequestIDKey).(string)
	if !ok {
		return nil, apperror.ErrMissingContext
	}

	runRepo := repo.NewRunRepo(ctx, username, requestID, dbPool)

	return service.NewRunService(ctx, runRepo, cfg), nil
}

func NewRunLogService(ctx context.Context) (*service.RunLogService, error) {
	cfg, dbPool, username, err := contextDefaultData(ctx)
	if err != nil {
		return nil, apperror.ErrMissingContext
	}

	runLogRepo := repo.NewRunLogRepo(ctx, username, dbPool)

	return service.NewRunLogService(ctx, runLogRepo, cfg), nil
}

func NewWorkflowEngineService(ctx context.Context) (*service.WorkflowEngineService, error) {
	cfg, ok := ctxkeys.GetValue(ctx, ctxkeys.AppConfigKey).(config.Configuration)
	if !ok {
		return nil, apperror.ErrMissingContext
	}

	return service.NewWorkflowEngineService(ctx, cfg), nil
}

func NewWorker(ctx context.Context, workerName string) (broker.BrokerAdapter, error) {
	switch workerName {
	// case "loopback":
	// 	return loopback.NewLoopback(ctx), nil
	default:
		return nil, apperror.NewGeneric(fmt.Errorf("worker adapter %s not found", workerName))
	}
}
