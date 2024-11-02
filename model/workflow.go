package model

import "time"

type Workflow struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password,omitempty"`
	Active   bool     `json:"active"`
	Meta     MetaData `json:"meta"`
}
type NewWorkflow struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	FailOnError bool         `json:"fail_on_error"`
	ObjectType  WorkflowType `json:"object_type"`
	Children    []Workflow   `json:"children"`
}

type WorkflowType string

const (
	Javascript WorkflowType = "javascript"
	External   WorkflowType = "external"
	Folder     WorkflowType = "folder"
)

type NewRun struct {
	WorkflowID string
	Error      error
	StartTime  time.Time
	FinishTime time.Time
}

type Run struct {
	WorkflowID string
	Error      error
	StartTime  time.Time
	FinishTime time.Time
	RunTime    string
	Meta       MetaData `json:"meta"`
}
