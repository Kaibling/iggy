package service

import (
	"context"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type workflowRepo interface {
	SaveWorkflow(newModel entity.NewWorkflow) (*entity.Workflow, error)
	FetchWorkflow(id string) (*entity.Workflow, error)
	FetchAll() ([]*entity.Workflow, error)
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

func (ts *WorkflowService) FetchWorkflow(id string) (*entity.Workflow, error) {
	return ts.repo.FetchWorkflow(id)
}

func (ts *WorkflowService) CreateWorkflow(u entity.NewWorkflow) (*entity.Workflow, error) {
	if u.ID == "" {
		u.ID = utils.NewULID().String()
	}
	return ts.repo.SaveWorkflow(u)
}

func (ts *WorkflowService) FetchAll() ([]*entity.Workflow, error) {
	return ts.repo.FetchAll()
}

func (ts *WorkflowService) DeleteWorkflow(id string) error {
	return ts.repo.DeleteWorkflow(id)
}
