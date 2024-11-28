package bootstrap

import (
	"context"

	"github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/iggy/entity"
	repo "github.com/kaibling/iggy/persistence/repository"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/service"
)

func NewWorkflowEngineService(ctx context.Context, cfg config.Configuration, dynDataService *service.DynDataService) (*service.WorkflowEngineService, error) { //nolint:lll
	return service.NewWorkflowEngineService(ctx, cfg, dynDataService), nil
}

func NewDynTabService(ctx context.Context, serviceConfig service.Config) (*service.DynTabService, error) {
	dynTabRepo := repo.NewDynTabRepo(ctx, serviceConfig.Username, serviceConfig.DBPool)

	return service.NewDynTabService(ctx, dynTabRepo, serviceConfig.Config), nil
}

func NewDynFieldService(ctx context.Context, serviceConfig service.Config) (*service.DynFieldService, error) {
	dynFieldRepo := repo.NewDynFieldRepo(ctx, serviceConfig.Username, serviceConfig.DBPool)

	dss, err := NewDynSchemaService(ctx, serviceConfig)
	if err != nil {
		return nil, err
	}

	return service.NewDynFieldService(ctx, dynFieldRepo, dss, serviceConfig.Config), nil
}

func NewDynDataService(ctx context.Context, serviceConfig service.Config) (*service.DynDataService, error) {
	dynTabRepo := repo.NewDynDataRepo(ctx, serviceConfig.Username, serviceConfig.DBPool)

	dss, err := NewDynSchemaService(ctx, serviceConfig)
	if err != nil {
		return nil, err
	}

	dfs, err := NewDynFieldService(ctx, serviceConfig)
	if err != nil {
		return nil, err
	}

	return service.NewDynDataService(ctx, dynTabRepo, dss, dfs, serviceConfig.Config), nil
}

func NewDynSchemaService(ctx context.Context, serviceConfig service.Config) (*service.DynSchemaService, error) {
	dynSchemaRepo := repo.NewDynSchemaRepo(ctx, serviceConfig.Username, serviceConfig.DBPool)

	return service.NewDynSchemaService(ctx, dynSchemaRepo, serviceConfig.Config), nil
}

func NewRunLogService(ctx context.Context, serviceConfig service.Config) (*service.RunLogService, error) {
	runLogRepo := repo.NewRunLogRepo(ctx, serviceConfig.Username, serviceConfig.DBPool)

	return service.NewRunLogService(ctx, runLogRepo, serviceConfig.Config), nil
}

func NewRunService(ctx context.Context, serviceConfig service.Config, task entity.Task) (*service.RunService, error) { //nolint:lll
	runRepo := repo.NewRunRepo(ctx, serviceConfig.Username, task.RequestID, serviceConfig.DBPool)

	return service.NewRunService(ctx, runRepo, task.UserID, serviceConfig.Config), nil
}

func NewWorkflowService(ctx context.Context, serviceConfig service.Config, scope string) (*service.WorkflowService, error) { //nolint:lll
	workflowRepo := repo.NewWorkflowRepo(ctx, serviceConfig.Username, serviceConfig.DBPool)

	return service.NewWorkflowService(ctx, workflowRepo, serviceConfig.Log.NewScope(scope), serviceConfig.Config), nil
}

func WorkerExecution(ctx context.Context, serviceConfig service.Config, task entity.Task) error {
	errs := apierror.NewMultiError()

	wfs, err := NewWorkflowService(ctx, serviceConfig, "worker_execution")
	errs.Add(err)

	rs, err := NewRunService(ctx, serviceConfig, task)
	errs.Add(err)

	rls, err := NewRunLogService(ctx, serviceConfig)
	errs.Add(err)

	dds, err := NewDynDataService(ctx, serviceConfig)
	errs.Add(err)

	wfes, err := NewWorkflowEngineService(ctx, serviceConfig.Config, dds)
	errs.Add(err)

	if errs.HasError() {
		return errs.GetErrors()[0]
	}

	if err := wfs.Execute(task.WorkflowID, wfes, rs, rls); err != nil {
		return err
	}

	return nil
}
