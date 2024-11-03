package service

import (
	"context"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type runLogRepo interface {
	SaveRunLog(newModel entity.NewRunLog) (*entity.RunLog, error)
	FetchRunLog(id string) (*entity.RunLog, error)
	FetchRunLogsByRun(runID string) ([]entity.RunLog, error)
}

type RunLogService struct {
	ctx  context.Context
	repo runLogRepo
	cfg  config.Configuration
}

func NewRunLogService(ctx context.Context, u runLogRepo, cfg config.Configuration) *RunLogService {
	return &RunLogService{ctx: ctx, repo: u, cfg: cfg}
}

func (ts *RunLogService) FetchRunLog(id string) (*entity.RunLog, error) {
	return ts.repo.FetchRunLog(id)
}

func (ts *RunLogService) CreateRunLog(u entity.NewRunLog) (*entity.RunLog, error) {
	if u.ID == "" {
		u.ID = utils.NewULID().String()
	}
	return ts.repo.SaveRunLog(u)
}

func (ts *RunLogService) FetchRunLogsByRun(runID string) ([]entity.RunLog, error) {
	return ts.repo.FetchRunLogsByRun(runID)
}
