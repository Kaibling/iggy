package service

import (
	"context"

	"github.com/kaibling/iggy/model"
	"github.com/kaibling/iggy/pkg/config"
)

type runRepo interface {
	SaveRun(newModel model.NewRun) (*model.Run, error)
	FetchRun(id string) (*model.Run, error)
	FetchRunByWorkflow(workflowID string) ([]*model.Run, error)
}

type RunService struct {
	ctx  context.Context
	repo runRepo
	cfg  config.Configuration
}

func NewRunService(ctx context.Context, u runRepo, cfg config.Configuration) *RunService {
	return &RunService{ctx: ctx, repo: u, cfg: cfg}
}

func (ts *RunService) FetchRun(id string) (*model.Run, error) {
	return ts.repo.FetchRun(id)
}

func (ts *RunService) CreateRun(u model.NewRun) (*model.Run, error) {
	return ts.repo.SaveRun(u)
}

func (ts *RunService) FetchRunByWorkflow(id string) ([]*model.Run, error) {
	return ts.repo.FetchRunByWorkflow(id)
}
