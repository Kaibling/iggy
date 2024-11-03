package workflow

import (
	"time"

	"github.com/kaibling/iggy/entity"
)

type WorkflowType string

const (
	Javascript WorkflowType = "javascript"
	External   WorkflowType = "external"
	Folder     WorkflowType = "folder"
)

type Workflow struct {
	ID          string
	Name        string
	Code        string
	ObjectType  WorkflowType
	Children    []Workflow
	FailOnError bool
}

func FromWorkflowEntity(e entity.Workflow) Workflow {
	children := []Workflow{}
	for _, c := range e.Children {
		children = append(children, FromWorkflowEntity(c))
	}
	return Workflow{
		ID:          e.ID,
		Name:        e.Name,
		Code:        e.Code,
		ObjectType:  WorkflowType(e.ObjectType),
		Children:    children,
		FailOnError: e.FailOnError,
	}
}

type Run struct {
	WorkflowID string
	Error      error
	SharedData map[string]any
	StartTime  time.Time
	FinishTime time.Time
}

func (r Run) ToNewEntity() entity.NewRun {
	newRun := entity.NewRun{
		WorkflowID: r.WorkflowID,
		StartTime:  r.StartTime,
		FinishTime: r.FinishTime,
	}
	if r.Error != nil {
		e := r.Error.Error()
		newRun.Error = &e
	}
	return newRun
}
