package service

import (
	"context"

	"github.com/kaibling/iggy/model"
	"github.com/kaibling/iggy/pkg/config"
)

type runLogRepo interface {
	SaveRunLog(newModel model.NewRunLog) (*model.RunLog, error)
	FetchRunLog(id string) (*model.RunLog, error)
	FetchRunLogsByRun(runID string) ([]model.RunLog, error)
}

type RunLogService struct {
	ctx  context.Context
	repo runLogRepo
	cfg  config.Configuration
}

func NewRunLogService(ctx context.Context, u runLogRepo, cfg config.Configuration) *RunLogService {
	return &RunLogService{ctx: ctx, repo: u, cfg: cfg}
}

func (ts *RunLogService) FetchRunLog(id string) (*model.RunLog, error) {
	return ts.repo.FetchRunLog(id)
}

func (ts *RunLogService) CreateRunLog(u model.NewRunLog) (*model.RunLog, error) {
	return ts.repo.SaveRunLog(u)
}

func (ts *RunLogService) FetchRunLogsByRun(runID string) ([]model.RunLog, error) {
	return ts.repo.FetchRunLogsByRun(runID)
}
