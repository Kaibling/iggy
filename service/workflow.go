package service

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/kaibling/apiforge/lib/utils"
	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
	"github.com/kaibling/iggy/pkg/git"
	"gopkg.in/yaml.v3"
)

type workflowRepo interface {
	CreateWorkflows(newModel []*entity.NewWorkflow) ([]entity.Workflow, error)
	UpdateWorkflow(id string, updateEntity entity.UpdateWorkflow) (*entity.Workflow, error)
	FetchToBackup() ([]entity.Workflow, error)
	FetchBackupAll() ([]entity.Workflow, error)
	FetchWorkflows(ids []string, depth int) ([]entity.Workflow, error)
	IDQuery(query string) ([]string, error)
	DeleteWorkflow(id string) error
	Upserts(workflows []entity.Workflow) error
}

type WorkflowService struct {
	ctx    context.Context
	repo   workflowRepo
	logger logging.Writer
	cfg    config.Configuration
}

func NewWorkflowService(ctx context.Context, u workflowRepo, logger logging.Writer, cfg config.Configuration) *WorkflowService { //nolint:lll, nolintlint
	return &WorkflowService{ctx: ctx, repo: u, logger: logger, cfg: cfg}
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

	if len(wf) == 0 {
		return fmt.Errorf("workflow not found: %s", workflowID)
	}

	// execute
	executedRuns := workflowExecutionService.Execute(wf[0])

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

func (ws *WorkflowService) ExportToGit(localPath, gitToken string) error {
	// export from DB
	exports, err := ws.repo.FetchToBackup()
	if err != nil {
		return err
	}

	// create files
	filenames, err := exportData(exports, localPath, ws.logger)
	if err != nil {
		return err
	}

	if err := git.CommitFiles(localPath, filenames, ws.logger); err != nil {
		return err
	}

	// push
	if err := git.Push(localPath, gitToken); err != nil {
		return err
	}

	return nil
}

func (ws *WorkflowService) ExportToDir(localPath string) error {
	// export from DB
	exports, err := ws.repo.FetchBackupAll()
	if err != nil {
		return err
	}
	// create files
	_, err = exportData(exports, localPath, ws.logger)
	if err != nil {
		return err
	}

	return nil
}

func (ws *WorkflowService) ImportFromFiles(localPath string) error {
	workflows, err := parseFiles(localPath)
	if err != nil {
		return err
	}

	fmt.Println(workflows)

	// update or create workflow
	//return ws.repo.Upserts(workflows)
	return nil
}

func exportData(data []entity.Workflow, path string, logger logging.Writer) ([]string, error) {
	fileList := []string{}

	b, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}

	filename := "export.yaml"
	//fileList = append(fileList, filename)

	absolutFilename := path + "/" + filename
	logger.Info("exporting " + absolutFilename)

	if err := os.WriteFile(absolutFilename, b, 0o600); err != nil { //nolint: mnd
		return nil, err
	}

	return fileList, nil
}

func parseFiles(localPath string) ([]entity.Workflow, error) {
	workflows := []entity.Workflow{}

	files, err := os.ReadDir(localPath)
	if err != nil {
		return nil, err
	}
	fmt.Println(localPath)
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}

		absulteFilePath := localPath + "/" + file.Name()

		dat, err := os.ReadFile(absulteFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed reading %s: %s", file.Name(), err)
		}
		var tmpWorkflow []entity.Workflow
		if err := yaml.Unmarshal(dat, &tmpWorkflow); err != nil {
			return nil, fmt.Errorf("failed unmarshaling file '%s': %s", file.Name(), err)
		}

		workflows = append(workflows, tmpWorkflow...)
	}

	return workflows, nil
}
