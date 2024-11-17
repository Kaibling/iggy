package service

import (
	"context"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

type workflowRepo interface {
	CreateWorkflows(newModel []*entity.NewWorkflow) ([]entity.Workflow, error)
	UpdateWorkflow(id string, updateEntity entity.UpdateWorkflow) (*entity.Workflow, error)
	// FetchWorkflow(id string) (*entity.Workflow, error)
	FetchWorkflows(ids []string, depth int) ([]entity.Workflow, error)
	IDQuery(query string) ([]string, error)
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

func (ws *WorkflowService) FetchWorkflows(ids []string) ([]entity.Workflow, error) {
	maxDepth := 10
	return ws.repo.FetchWorkflows(ids, maxDepth)

}

func (ws *WorkflowService) CreateWorkflows(newWorkflows []*entity.NewWorkflow) ([]entity.Workflow, error) {
	for _, wf := range newWorkflows {
		if wf.ID == "" {
			wf.ID = utils.NewULID().String()
		}
	}

	return ws.repo.CreateWorkflows(newWorkflows)
}

func (ws *WorkflowService) Update(id string, updateEntity entity.UpdateWorkflow) (*entity.Workflow, error) {
	return ws.repo.UpdateWorkflow(id, updateEntity)
}

func (ws *WorkflowService) DeleteWorkflow(id string) error {
	return ws.repo.DeleteWorkflow(id)
}

func (ws *WorkflowService) Execute(workflowID string, workflowExecutionService *WorkflowEngineService, runService *RunService, runLogService *RunLogService) error { //nolint: lll
	// get workflow
	wf, err := ws.FetchWorkflows([]string{workflowID})
	if err != nil {
		return err
	}

	// execute
	executedRuns, err := workflowExecutionService.Execute(wf[0])
	if err != nil {
		return err
	}

	// save runs
	for _, newRun := range executedRuns {
		if err := runService.CreateRun(newRun, runLogService); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WorkflowService) FetchByPagination(qp params.Pagination) ([]entity.Workflow, params.Pagination, error) {
	p := NewPagination(qp, "workflows")

	idQuery := p.GetCursorSQL()

	ids, err := ws.repo.IDQuery(idQuery)
	if err != nil {
		return nil, params.Pagination{}, err
	}

	ids, pag := p.FinishPagination(ids)

	loadedWorkflows, err := ws.repo.FetchWorkflows(ids, 1)
	if err != nil {
		return nil, params.Pagination{}, err
	}

	return loadedWorkflows, pag, nil
}
