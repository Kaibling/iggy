package workflow

import (
	"time"

	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/workflow/adapter"
)

type WorkflowAdapter interface {
	Execute(code string, sharedData map[string]any) adapter.ExecutionResult
}

type WorkflowResult struct {
	Error      error
	Runs       []Run
	SharedData map[string]any
}

type Engine struct{}

func NewEngine() Engine {
	return Engine{}
}

func (e Engine) Execute(workflow entity.Workflow) ([]entity.NewRun, error) {
	wf := FromWorkflowEntity(workflow)
	sharedData := map[string]any{}
	result := e.execute_workflow(wf, sharedData)
	runs := []entity.NewRun{}
	for _, r := range result.Runs {
		runs = append(runs, r.ToNewEntity())
	}
	return runs, result.Error

}

func (e Engine) execute_workflow(w Workflow, sharedData map[string]any) WorkflowResult {
	if w.ObjectType == Folder {
		return e.execute_folder(w, sharedData)
	}
	return e.execute_adapter(w, sharedData)
}

func (e Engine) execute_folder(w Workflow, sharedData map[string]any) WorkflowResult {
	runs := []Run{}
	for _, child := range w.Children {
		result := e.execute_workflow(child, sharedData)
		runs = append(runs, result.Runs...)
		sharedData = result.SharedData

		// handle fail on error
		if result.Error != nil && child.FailOnError {
			return WorkflowResult{
				Error: nil,
				Runs:  runs,
			}
		}

	}
	return WorkflowResult{
		Error: nil,
		Runs:  runs,
	}
}

func (e Engine) execute_adapter(w Workflow, sharedData map[string]any) WorkflowResult {
	execAdapter, err := e.getAdapter(w.ObjectType)
	if err != nil {
		return WorkflowResult{
			Error:      err,
			SharedData: sharedData,
		}
	}

	startTime := time.Now()
	result := execAdapter.Execute(w.Code, sharedData)
	finishTime := time.Now()

	newRun := Run{
		WorkflowID: w.ID,
		Error:      result.Error,
		SharedData: result.SharedData,
		StartTime:  startTime,
		FinishTime: finishTime,
		Logs:       result.Logs,
	}
	return WorkflowResult{
		Error:      result.Error,
		SharedData: result.SharedData,
		Runs:       []Run{newRun},
	}
}

func (e Engine) getAdapter(wft WorkflowType) (WorkflowAdapter, error) {
	switch wft {
	case Javascript:
		return adapter.NewJavascriptAdapter(), nil
	default:
		return nil, apperror.MissingJSAdapter
	}
}
