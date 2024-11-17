package service

import (
	"context"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type runRepo interface {
	SaveRun(newModel entity.NewRun) (*entity.Run, error)
	FetchRuns(id []string) ([]*entity.Run, error)
	FetchRunByWorkflow(workflowID string) ([]*entity.Run, error)
	FetchRunByRequestID(requestID string) (*entity.Run, error)
	IDQuery(query string) ([]string, error)
}

type RunService struct {
	ctx    context.Context
	repo   runRepo
	cfg    config.Configuration
	userID string
}

func NewRunService(ctx context.Context, repo runRepo, userID string, cfg config.Configuration) *RunService {
	return &RunService{ctx, repo, cfg, userID}
}

func (ts *RunService) FetchRuns(ids []string) ([]*entity.Run, error) {
	fetchedRuns, err := ts.repo.FetchRuns(ids)
	if err != nil {
		return nil, err
	}

	for _, run := range fetchedRuns {
		run.CalculateRuntime()
	}

	return fetchedRuns, nil
}

func (ts *RunService) CreateRun(newEntity entity.NewRun, runLogService *RunLogService) error {
	if newEntity.ID == "" {
		newEntity.ID = utils.NewULID().String()
	}

	newEntity.UserID = ts.userID

	_, err := ts.repo.SaveRun(newEntity)
	if err != nil {
		return err
	}

	runLogs := []entity.NewRunLog{}

	for _, runLog := range newEntity.Logs {
		runLog.ID = utils.NewULID().String()
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

func (ts *RunService) FetchRunByRequestID(requestID string) (*entity.Run, error) {
	return ts.repo.FetchRunByRequestID(requestID)
}

func (ts *RunService) FetchByPagination(qp params.Pagination) ([]*entity.Run, params.Pagination, error) {
	p := NewPagination(qp, "runs")

	idQuery := p.GetCursorSQL()

	ids, err := ts.repo.IDQuery(idQuery)
	if err != nil {
		return nil, params.Pagination{}, err
	}

	ids, pag := p.FinishPagination(ids)

	loadedRuns, err := ts.FetchRuns(ids)
	if err != nil {
		return nil, params.Pagination{}, err
	}

	return loadedRuns, pag, nil
}
