package service

import (
	"context"

	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/pkg/workflow"
)

type WorkflowEngineService struct {
	ctx  context.Context
	repo workflow.Engine
	cfg  config.Configuration
}

func NewWorkflowEngineService(ctx context.Context, cfg config.Configuration, objService *DynDataService) *WorkflowEngineService { //nolint:lll
	return &WorkflowEngineService{ctx: ctx, repo: workflow.NewEngine(objService), cfg: cfg}
}

func (ts *WorkflowEngineService) Execute(w entity.Workflow) []entity.NewRun {
	return ts.repo.Execute(w)
}
