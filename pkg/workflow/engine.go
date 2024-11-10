package workflow

import (
	"time"

	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/workflow/adapter"
)

type Adapter interface {
	Execute(code string, sharedData map[string]any) adapter.ExecutionResult
}

type Result struct {
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
	result := e.executeWorkflow(wf, sharedData)
	runs := []entity.NewRun{}

	for _, r := range result.Runs {
		runs = append(runs, r.ToNewEntity())
	}

	return runs, result.Error
}

func (e Engine) executeWorkflow(w Workflow, sharedData map[string]any) Result {
	if w.ObjectType == Folder {
		return e.executeFolder(w, sharedData)
	}

	return e.executeAdapter(w, sharedData)
}

func (e Engine) executeFolder(w Workflow, sharedData map[string]any) Result {
	runs := []Run{}

	for _, child := range w.Children {
		result := e.executeWorkflow(child, sharedData)
		runs = append(runs, result.Runs...)
		sharedData = result.SharedData

		// handle fail on error
		if result.Error != nil && child.FailOnError {
			return Result{
				Error: nil,
				Runs:  runs,
			}
		}
	}

	return Result{
		Error: nil,
		Runs:  runs,
	}
}

func (e Engine) executeAdapter(w Workflow, sharedData map[string]any) Result {
	execAdapter, err := e.getAdapter(w.ObjectType)
	if err != nil {
		return Result{
			Error:      err,
			SharedData: sharedData,
			Runs:       []Run{},
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

	return Result{
		Error:      result.Error,
		SharedData: result.SharedData,
		Runs:       []Run{newRun},
	}
}

func (e Engine) getAdapter(wft Type) (Adapter, error) { //nolint: ireturn
	switch wft {
	case Javascript:
		return adapter.NewJavascriptAdapter(), nil
	case External:
		panic("not implemented")
	case Folder:
		panic("not implemented")
	default:
		return nil, apperror.ErrMissingJSAdapter
	}
}
