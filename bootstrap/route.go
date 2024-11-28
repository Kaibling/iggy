package bootstrap

import (
	"context"

	"github.com/kaibling/apiforge/ctxkeys"
	repo "github.com/kaibling/iggy/persistence/repository"
	"github.com/kaibling/iggy/service"
)

func BuildRouteUserService(ctx context.Context) (*service.UserService, error) {
	cfg, dbPool, _, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	userRepo := repo.NewUserRepo(ctx, username, dbPool)

	return service.NewUserService(ctx, userRepo, cfg), nil
}

func BuildRouteTokenService(ctx context.Context) (*service.TokenService, error) {
	cfg, dbPool, _, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	tokenRepo := repo.NewTokenRepo(ctx, username, dbPool)

	return service.NewTokenService(ctx, tokenRepo, cfg), nil
}

func NewUserServiceAnonym(ctx context.Context, username string) (*service.UserService, error) {
	cfg, dbPool, _, err := contextBasisData(ctx)
	if err != nil {
		return nil, err
	}

	userRepo := repo.NewUserRepo(ctx, username, dbPool)

	return service.NewUserService(ctx, userRepo, cfg), nil
}

func NewTokenServiceAnonym(ctx context.Context, username string) (*service.TokenService, error) {
	cfg, dbPool, _, err := contextBasisData(ctx)
	if err != nil {
		return nil, err
	}

	tokenRepo := repo.NewTokenRepo(ctx, username, dbPool)

	return service.NewTokenService(ctx, tokenRepo, cfg), nil
}

func BuildRouteWorkflowService(ctx context.Context) (*service.WorkflowService, error) {
	cfg, dbPool, log, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	workflowRepo := repo.NewWorkflowRepo(ctx, username, dbPool)

	return service.NewWorkflowService(ctx, workflowRepo, log, cfg), nil
}

func BuildRouteRunService(ctx context.Context) (*service.RunService, error) {
	cfg, dbPool, _, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	requestID, ok := ctxkeys.GetValue(ctx, ctxkeys.RequestIDKey).(string)
	if !ok {
		return nil, err
	}

	userID, ok := ctxkeys.GetValue(ctx, ctxkeys.UserIDKey).(string)
	if !ok {
		return nil, err
	}

	runRepo := repo.NewRunRepo(ctx, username, requestID, dbPool)

	return service.NewRunService(ctx, runRepo, userID, cfg), nil
}

func BuildRouteDynTabService(ctx context.Context) (*service.DynTabService, error) {
	cfg, dbPool, _, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	dynTabRepo := repo.NewDynTabRepo(ctx, username, dbPool)

	return service.NewDynTabService(ctx, dynTabRepo, cfg), nil
}

func BuildRouteRunLogService(ctx context.Context) (*service.RunLogService, error) {
	cfg, dbPool, _, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	runLogRepo := repo.NewRunLogRepo(ctx, username, dbPool)

	return service.NewRunLogService(ctx, runLogRepo, cfg), nil
}

func BuildRouteDynSchemaService(ctx context.Context) (*service.DynSchemaService, error) {
	cfg, dbPool, _, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	dynSchemaRepo := repo.NewDynSchemaRepo(ctx, username, dbPool)

	return service.NewDynSchemaService(ctx, dynSchemaRepo, cfg), nil
}

func BuildRouteDynFieldService(ctx context.Context) (*service.DynFieldService, error) {
	cfg, dbPool, _, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	dss, err := BuildRouteDynSchemaService(ctx)
	if err != nil {
		return nil, err
	}

	dynFieldRepo := repo.NewDynFieldRepo(ctx, username, dbPool)

	return service.NewDynFieldService(ctx, dynFieldRepo, dss, cfg), nil
}

// func NewWorkflowEngineService(ctx context.Context) (*service.WorkflowEngineService, error) {
// 	cfg, ok := ctxkeys.GetValue(ctx, ctxkeys.AppConfigKey).(config.Configuration)
// 	if !ok {
// 		return nil, apperror.ErrNewMissingContext("config")
// 	}

// 	dynService, err := NewDynTabService(ctx)
// 	if err != nil {
// 		return nil, apperror.ErrMissing
// 	}

// 	return service.NewWorkflowEngineService(ctx, cfg, dynService), nil
// }
