package bootstrap

import (
	"context"

	"github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/entity"
	repo "github.com/kaibling/iggy/persistence/repository"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/service"
)

func NewWorkflowEngineServiceV2(ctx context.Context, cfg config.Configuration, dynDataService *service.DynDataService) (*service.WorkflowEngineService, error) { //nolint:lll
	return service.NewWorkflowEngineService(ctx, cfg, dynDataService), nil
}

func NewDynTabServiceV2(ctx context.Context, brokerConfig broker.SubscriberConfig) (*service.DynTabService, error) {
	dynTabRepo := repo.NewDynTabRepo(ctx, brokerConfig.Username, brokerConfig.DBPool)

	return service.NewDynTabService(ctx, dynTabRepo, brokerConfig.Config), nil
}

func NewDynFieldServiceV2(ctx context.Context, brokerConfig broker.SubscriberConfig) (*service.DynFieldService, error) {
	dynFieldRepo := repo.NewDynFieldRepo(ctx, brokerConfig.Username, brokerConfig.DBPool)

	dss, err := NewDynSchemaServiceV2(ctx, brokerConfig)
	if err != nil {
		return nil, err
	}

	return service.NewDynFieldService(ctx, dynFieldRepo, dss, brokerConfig.Config), nil
}

func NewDynDataServiceV2(ctx context.Context, brokerConfig broker.SubscriberConfig) (*service.DynDataService, error) {
	dynTabRepo := repo.NewDynDataRepo(ctx, brokerConfig.Username, brokerConfig.DBPool)

	dss, err := NewDynSchemaServiceV2(ctx, brokerConfig)
	if err != nil {
		return nil, err
	}

	dfs, err := NewDynFieldServiceV2(ctx, brokerConfig)
	if err != nil {
		return nil, err
	}

	return service.NewDynDataService(ctx, dynTabRepo, dss, dfs, brokerConfig.Config), nil
}

func NewDynSchemaServiceV2(ctx context.Context, brokerConfig broker.SubscriberConfig) (*service.DynSchemaService, error) {
	dynSchemaRepo := repo.NewDynSchemaRepo(ctx, brokerConfig.Username, brokerConfig.DBPool)

	return service.NewDynSchemaService(ctx, dynSchemaRepo, brokerConfig.Config), nil
}

func NewRunLogServiceV2(ctx context.Context, brokerConfig broker.SubscriberConfig) (*service.RunLogService, error) {
	runLogRepo := repo.NewRunLogRepo(ctx, brokerConfig.Username, brokerConfig.DBPool)

	return service.NewRunLogService(ctx, runLogRepo, brokerConfig.Config), nil
}

func NewRunServiceV2(ctx context.Context, brokerConfig broker.SubscriberConfig, task entity.Task) (*service.RunService, error) { //nolint:lll
	runRepo := repo.NewRunRepo(ctx, brokerConfig.Username, task.RequestID, brokerConfig.DBPool)

	return service.NewRunService(ctx, runRepo, task.UserID, brokerConfig.Config), nil
}

func NewWorkflowServiceV2(ctx context.Context, brokerConfig broker.SubscriberConfig) (*service.WorkflowService, error) {
	workflowRepo := repo.NewWorkflowRepo(ctx, brokerConfig.Username, brokerConfig.DBPool)

	return service.NewWorkflowService(ctx, workflowRepo, brokerConfig.Log), nil
}

func WorkerExecution(ctx context.Context, brokerConfig broker.SubscriberConfig, task entity.Task) error {
	errs := apierror.NewMultiError()

	wfs, err := NewWorkflowServiceV2(ctx, brokerConfig)
	errs.Add(err)

	rs, err := NewRunServiceV2(ctx, brokerConfig, task)
	errs.Add(err)

	rls, err := NewRunLogServiceV2(ctx, brokerConfig)
	errs.Add(err)

	// dts, err := NewDynTabServiceV2(ctx, brokerConfig)
	// errs.Add(err)

	dds, err := NewDynDataServiceV2(ctx, brokerConfig)
	errs.Add(err)

	wfes, err := NewWorkflowEngineServiceV2(ctx, brokerConfig.Config, dds)
	errs.Add(err)

	if errs.HasError() {
		return errs.GetErrors()[0]
	}

	if err := wfs.Execute(task.WorkflowID, wfes, rs, rls); err != nil {
		return err
	}

	return nil
}
