package service

import (
	"context"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type runRepo interface {
	SaveRun(newModel entity.NewRun) (*entity.Run, error)
	FetchRun(id string) (*entity.Run, error)
	FetchRunByWorkflow(workflowID string) ([]*entity.Run, error)
}

type RunService struct {
	ctx  context.Context
	repo runRepo
	cfg  config.Configuration
}

func NewRunService(ctx context.Context, u runRepo, cfg config.Configuration) *RunService {
	return &RunService{ctx: ctx, repo: u, cfg: cfg}
}

func (ts *RunService) FetchRun(id string) (*entity.Run, error) {
	return ts.repo.FetchRun(id)
}

func (ts *RunService) CreateRun(newEntity entity.NewRun, runLogService *RunLogService) error {
	if newEntity.ID == "" {
		newEntity.ID = utils.NewULID().String()
	}

	_, err := ts.repo.SaveRun(newEntity)
	if err != nil {
		return err
	}

	runLogs := []entity.NewRunLog{}

	for _, runLog := range newEntity.Logs {
		runLog.RunID = newEntity.ID
		runLogs = append(runLogs, runLog)
	}

	if err := runLogService.CreateRunLogs(runLogs); err != nil {
		return err
	}
	return nil
}

func (ts *RunService) FetchRunByWorkflow(id string) ([]*entity.Run, error) {
	return ts.repo.FetchRunByWorkflow(id)
}
