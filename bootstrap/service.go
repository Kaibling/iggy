package bootstrap

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/entity"
	repo "github.com/kaibling/iggy/persistence/repository"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/service"
)

func ContextDefaultData(ctx context.Context) (config.Configuration, *pgxpool.Pool, string, error) {
	cfg, dbPool, err := contextBasisData(ctx)
	if err != nil {
		return cfg, dbPool, "", err
	}

	username, ok := ctxkeys.GetValue(ctx, ctxkeys.UserNameKey).(string)
	if !ok {
		return cfg, dbPool, "", apperror.ErrNewMissingContext("username")
	}

	return cfg, dbPool, username, nil
}

func contextBasisData(ctx context.Context) (config.Configuration, *pgxpool.Pool, error) {
	cfg, ok := ctxkeys.GetValue(ctx, ctxkeys.AppConfigKey).(config.Configuration)
	if !ok {
		return config.Configuration{}, nil, apperror.ErrNewMissingContext("config")
	}

	dbPool, ok := ctxkeys.GetValue(ctx, ctxkeys.DBConnKey).(*pgxpool.Pool)
	if !ok {
		return config.Configuration{}, nil, apperror.ErrNewMissingContext("db pool")
	}

	return cfg, dbPool, nil
}

func NewUserService(ctx context.Context) (*service.UserService, error) {
	cfg, dbPool, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	userRepo := repo.NewUserRepo(ctx, username, dbPool)

	return service.NewUserService(ctx, userRepo, cfg), nil
}

func NewTokenService(ctx context.Context) (*service.TokenService, error) {
	cfg, dbPool, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	tokenRepo := repo.NewTokenRepo(ctx, username, dbPool)

	return service.NewTokenService(ctx, tokenRepo, cfg), nil
}

func NewUserServiceAnonym(ctx context.Context, username string) (*service.UserService, error) {
	cfg, dbPool, err := contextBasisData(ctx)
	if err != nil {
		return nil, err
	}

	userRepo := repo.NewUserRepo(ctx, username, dbPool)

	return service.NewUserService(ctx, userRepo, cfg), nil
}

func NewTokenServiceAnonym(ctx context.Context, username string) (*service.TokenService, error) {
	cfg, dbPool, err := contextBasisData(ctx)
	if err != nil {
		return nil, err
	}

	tokenRepo := repo.NewTokenRepo(ctx, username, dbPool)

	return service.NewTokenService(ctx, tokenRepo, cfg), nil
}

func NewWorkflowService(ctx context.Context) (*service.WorkflowService, error) {
	cfg, dbPool, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	workflowRepo := repo.NewWorkflowRepo(ctx, username, dbPool)

	return service.NewWorkflowService(ctx, workflowRepo, cfg), nil
}

func NewWorkflowServiceV2(ctx context.Context, cfg config.Configuration, dbPool *pgxpool.Pool, username string) (*service.WorkflowService, error) { //nolint: lll
	workflowRepo := repo.NewWorkflowRepo(ctx, username, dbPool)

	return service.NewWorkflowService(ctx, workflowRepo, cfg), nil
}

func NewRunService(ctx context.Context) (*service.RunService, error) {
	cfg, dbPool, username, err := ContextDefaultData(ctx)
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

func NewDynTabService(ctx context.Context) (*service.DynTabService, error) {
	cfg, dbPool, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	dynTabRepo := repo.NewDynTabRepo(ctx, username, dbPool)

	return service.NewDynTabService(ctx, dynTabRepo, cfg), nil
}

func NewRunServiceV2(ctx context.Context,
	cfg config.Configuration,
	dbPool *pgxpool.Pool,
	username,
	userID string,
	requestID string,
) (*service.RunService, error) { //nolint: lll,nolintlint
	runRepo := repo.NewRunRepo(ctx, username, requestID, dbPool)

	return service.NewRunService(ctx, runRepo, userID, cfg), nil
}

func NewRunLogService(ctx context.Context) (*service.RunLogService, error) {
	cfg, dbPool, username, err := ContextDefaultData(ctx)
	if err != nil {
		return nil, err
	}

	runLogRepo := repo.NewRunLogRepo(ctx, username, dbPool)

	return service.NewRunLogService(ctx, runLogRepo, cfg), nil
}

func NewRunLogServiceV2(ctx context.Context, cfg config.Configuration, dbPool *pgxpool.Pool, username string) (*service.RunLogService, error) { //nolint: lll
	runLogRepo := repo.NewRunLogRepo(ctx, username, dbPool)

	return service.NewRunLogService(ctx, runLogRepo, cfg), nil
}

func NewWorkflowEngineService(ctx context.Context) (*service.WorkflowEngineService, error) {
	cfg, ok := ctxkeys.GetValue(ctx, ctxkeys.AppConfigKey).(config.Configuration)
	if !ok {
		return nil, apperror.ErrNewMissingContext("config")
	}

	return service.NewWorkflowEngineService(ctx, cfg), nil
}

func NewWorkflowEngineServiceV2(ctx context.Context, cfg config.Configuration) (*service.WorkflowEngineService, error) {
	return service.NewWorkflowEngineService(ctx, cfg), nil
}

func ContextParams(ctx context.Context) (params.Pagination, error) {
	p, ok := ctxkeys.GetValue(ctx, ctxkeys.PaginationKey).(params.Pagination)
	if !ok {
		return params.Pagination{}, apperror.ErrNewMissingContext("pagination")
	}

	return p, nil
}

func ContextToken(ctx context.Context) (string, error) {
	t, ok := ctxkeys.GetValue(ctx, ctxkeys.TokenKey).(string)
	if !ok {
		return "", apperror.ErrNewMissingContext("token")
	}

	return t, nil
}

func ContextLogger(ctx context.Context) (logging.Writer, error) { //nolint: ireturn
	l, ok := ctxkeys.GetValue(ctx, ctxkeys.LoggerKey).(logging.Writer)
	if !ok {
		return nil, apperror.ErrNewMissingContext("logger")
	}

	return l, nil
}

func ContextRequestID(ctx context.Context) (string, error) {
	requestID, ok := ctxkeys.GetValue(ctx, ctxkeys.RequestIDKey).(string)
	if !ok {
		return "", apperror.ErrNewMissingContext("requestID")
	}

	return requestID, nil
}

func ContextUserID(ctx context.Context) (string, error) {
	userID, ok := ctxkeys.GetValue(ctx, ctxkeys.UserIDKey).(string)
	if !ok {
		return "", apperror.ErrNewMissingContext("userID")
	}

	return userID, nil
}

func WorkerExecution(ctx context.Context, brokerConfig broker.SubscriberConfig, task entity.Task) error {
	wfs, err := NewWorkflowServiceV2(ctx, brokerConfig.Config, brokerConfig.DBPool, brokerConfig.Username)
	if err != nil {
		return err
	}

	rs, err := NewRunServiceV2(ctx, brokerConfig.Config, brokerConfig.DBPool, brokerConfig.Username, task.UserID, task.RequestID) //nolint: lll
	if err != nil {
		return err
	}

	rls, err := NewRunLogServiceV2(ctx, brokerConfig.Config, brokerConfig.DBPool, brokerConfig.Username)
	if err != nil {
		return err
	}

	wfes, err := NewWorkflowEngineServiceV2(ctx, brokerConfig.Config)
	if err != nil {
		return err
	}

	if err := wfs.Execute(task.WorkflowID, wfes, rs, rls); err != nil {
		return err
	}

	return nil
}
