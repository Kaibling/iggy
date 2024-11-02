package workflow

import "time"

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

type Run struct {
	WorkflowID string
	Error      error
	SharedData map[string]any
	StartTime  time.Time
	FinishTime time.Time
}
