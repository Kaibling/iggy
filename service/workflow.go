package service

import (
	"context"

	"github.com/kaibling/iggy/model"
	"github.com/kaibling/iggy/pkg/config"
)

type workflowRepo interface {
	SaveWorkflow(newModel model.NewWorkflow) (*model.Workflow, error)
	FetchWorkflow(id string) (*model.Workflow, error)
	FetchAll() ([]*model.Workflow, error)
	DeleteWorkflow(id string) error
}

type WorkflowService struct {
	ctx  context.Context
	repo workflowRepo
	cfg  config.Configuration
}

func NewWorkflowService(ctx context.Context, u workflowRepo, cfg config.Configuration) *WorkflowService {
	return &WorkflowService{ctx: ctx, repo: u, cfg: cfg}
}

func (ts *WorkflowService) FetchWorkflow(id string) (*model.Workflow, error) {
	return ts.repo.FetchWorkflow(id)
}

func (ts *WorkflowService) CreateWorkflow(u model.NewWorkflow) (*model.Workflow, error) {
	return ts.repo.SaveWorkflow(u)
}

func (ts *WorkflowService) FetchAll() ([]*model.Workflow, error) {
	return ts.repo.FetchAll()
}

func (ts *WorkflowService) DeleteWorkflow(id string) error {
	return ts.repo.DeleteWorkflow(id)
}
