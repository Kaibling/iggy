package service

import (
	"context"
	"fmt"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/pkg/utility"
)

type workflowRepo interface {
	SaveWorkflow(newModel entity.NewWorkflow) (*entity.Workflow, error)
	UpdateWorkflow(id string, updateEntity entity.UpdateWorkflow) (*entity.Workflow, error)
	//FetchWorkflow(id string) (*entity.Workflow, error)
	FetchWorkflows(ids []string, depth int) ([]entity.Workflow, error)
	//FetchAll() ([]*entity.Workflow, error)
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

func (ws *WorkflowService) FetchWorkflows(ids []string) (*entity.Workflow, error) {
	maxDepth := 10
	if workflows, err := ws.repo.FetchWorkflows(ids, maxDepth); err != nil {
		return nil, err
	} else {
		return &workflows[0], nil
	}

}

func (ws *WorkflowService) CreateWorkflow(u entity.NewWorkflow) (*entity.Workflow, error) {
	if u.ID == "" {
		u.ID = utils.NewULID().String()
	}
	return ws.repo.SaveWorkflow(u)
}

func (ws *WorkflowService) Update(id string, updateEntity entity.UpdateWorkflow) (*entity.Workflow, error) {
	return ws.repo.UpdateWorkflow(id, updateEntity)
}

func (ws *WorkflowService) DeleteWorkflow(id string) error {
	return ws.repo.DeleteWorkflow(id)
}

func (ws *WorkflowService) Execute(workflowID string, workflowExecutionService *WorkflowEngineService, runService *RunService, runLogService *RunLogService) error {
	// get workflow
	wf, err := ws.FetchWorkflows([]string{workflowID})
	if err != nil {
		return err
	}

	// execute
	executedRuns, err := workflowExecutionService.Execute(*wf)
	if err != nil {
		return err
	}
	fmt.Println(utility.Pretty(executedRuns))
	// save runs
	for _, newRun := range executedRuns {
		if err := runService.CreateRun(newRun, runLogService); err != nil {
			return err
		}

	}

	return nil
}
