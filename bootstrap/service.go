package bootstrap

import (
	"context"
	"fmt"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/apperror"
	repo "github.com/kaibling/iggy/persistence/repository"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/service"
)

func NewUserService(ctx context.Context) *service.UserService {
	username := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)

	userRepo := repo.NewUserRepo(ctx, username)
	return service.NewUserService(ctx, userRepo, cfg)
}

func NewTokenService(ctx context.Context) *service.TokenService {
	username := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)

	tokenRepo := repo.NewTokenRepo(ctx, username)
	return service.NewTokenService(ctx, tokenRepo, cfg)
}

func NewUserServiceAnonym(ctx context.Context, username string) *service.UserService {
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)
	userRepo := repo.NewUserRepo(ctx, username)
	return service.NewUserService(ctx, userRepo, cfg)
}

func NewTokenServiceAnonym(ctx context.Context, username string) *service.TokenService {
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)
	tokenRepo := repo.NewTokenRepo(ctx, username)
	return service.NewTokenService(ctx, tokenRepo, cfg)
}

func NewWorkflowService(ctx context.Context) *service.WorkflowService {
	username := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)

	workflowRepo := repo.NewWorkflowRepo(ctx, username)
	return service.NewWorkflowService(ctx, workflowRepo, cfg)
}

func NewRunService(ctx context.Context) *service.RunService {
	username := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)
	requestID := ctxkeys.GetValue(ctx, ctxkeys.RequestIDKey).(string)

	runRepo := repo.NewRunRepo(ctx, username, requestID)
	return service.NewRunService(ctx, runRepo, cfg)
}

func NewRunLogService(ctx context.Context) *service.RunLogService {
	username := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)

	runLogRepo := repo.NewRunLogRepo(ctx, username)
	return service.NewRunLogService(ctx, runLogRepo, cfg)
}

func NewWorkflowEngineService(ctx context.Context) *service.WorkflowEngineService {
	cfg := ctxkeys.GetValue(ctx, "cfg").(config.Configuration)
	return service.NewWorkflowEngineService(ctx, cfg)
}

func NewWorker(ctx context.Context, workerName string) (broker.BrokerAdapter, error) {
	switch workerName {
	// case "loopback":
	// 	return loopback.NewLoopback(ctx), nil
	default:
		return nil, apperror.NewGeneric(fmt.Errorf("worker adapter %s not found", workerName))
	}
}
