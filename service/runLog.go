package service

import (
	"context"
	"time"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type runLogRepo interface {
	CreateRunLog(newModel entity.NewRunLog) (*entity.RunLog, error)
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

func (rls *RunLogService) FetchRunLog(id string) (*entity.RunLog, error) {
	return rls.repo.FetchRunLog(id)
}
func (rls *RunLogService) CreateRunLogs(newEntities []entity.NewRunLog) error {
	for _, e := range newEntities {
		if _, err := rls.repo.CreateRunLog(e); err != nil {
			return err
		}
	}
	return nil
}

func (rls *RunLogService) CreateRunLog(u entity.NewRunLog) (*entity.RunLog, error) {
	if u.ID == "" {
		u.ID = utils.NewULID().String()
	}
	return rls.repo.CreateRunLog(u)
}

func (rls *RunLogService) FetchRunLogsByRun(runID string) ([]entity.RunLog, error) {
	return rls.repo.FetchRunLogsByRun(runID)
}

func (rls *RunLogService) Log(msg, runID string) error {
	logLine := entity.NewRunLog{ //nolint: exhaustruct
		RunID:     runID,
		Message:   msg,
		Timestamp: time.Now(),
	}

	if _, err := rls.CreateRunLog(logLine); err != nil {
		return err
	}
	return nil
}
