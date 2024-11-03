package service

import (
	"context"
	"fmt"
	"time"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/pkg/utility"
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
func (ts *RunLogService) CreateRunLogs(newEntities []entity.NewRunLog) error {
	for _, e := range newEntities {
		if _, err := ts.repo.SaveRunLog(e); err != nil {
			return err
		}
	}
	return nil
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

func (s *RunLogService) Log(msg, runID string) error {
	logLine := entity.NewRunLog{
		RunID:     runID,
		Message:   msg,
		Timestamp: time.Now(),
	}
	fmt.Println(utility.Pretty(logLine))
	if _, err := s.CreateRunLog(logLine); err != nil {
		return err
	}
	return nil
}
