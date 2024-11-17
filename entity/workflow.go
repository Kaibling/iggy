package entity

import (
	"time"

	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/pkg/utility"
)

type Workflow struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	BuildIn     bool         `json:"build_in"`
	FailOnError bool         `json:"fail_on_error"`
	ObjectType  WorkflowType `json:"object_type"`
	Children    []Workflow   `json:"children"`
	Meta        MetaData     `json:"meta"`
}

type NewWorkflow struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Code        string       `json:"code"`
	BuildIn     bool         `json:"build_in"`
	FailOnError bool         `json:"fail_on_error"`
	ObjectType  WorkflowType `json:"object_type"`
	Children    []Identifier `json:"children"`
}

// TODO add more data, what is wrong.
func (w NewWorkflow) Validate() *apperror.AppError {
	if w.ObjectType == "" {
		return &apperror.ErrMalformedRequest
	}

	return nil
}

type UpdateWorkflow struct {
	Name        *string       `json:"name"`
	Code        *string       `json:"code"`
	BuildIn     *bool         `json:"build_in"`
	FailOnError *bool         `json:"fail_on_error"`
	ObjectType  *WorkflowType `json:"object_type"`
	Children    []Identifier  `json:"children"`
}

type WorkflowType string

const (
	Javascript WorkflowType = "javascript"
	External   WorkflowType = "external"
	Folder     WorkflowType = "folder"
)

type NewRun struct {
	ID         string
	WorkflowID string
	UserID     string
	Error      *string
	StartTime  time.Time
	FinishTime time.Time
	Logs       []NewRunLog
}

type Run struct {
	ID         string     `json:"id"`
	Workflow   Identifier `json:"workflow"`
	User       Identifier `json:"user"`
	Error      *string    `json:"error"`
	StartTime  time.Time  `json:"start_time"`
	FinishTime time.Time  `json:"finish_time"`
	RunTime    string     `json:"run_time"`
	Meta       MetaData   `json:"meta"`
}

func (r *Run) CalculateRuntime() {
	r.RunTime = utility.FormatDuration(r.FinishTime.Sub(r.StartTime))
}

type NewRunLog struct {
	ID        string
	RunID     string
	Message   string
	Timestamp time.Time
}

type RunLog struct {
	ID        string    `json:"id"`
	RunID     string    `json:"run_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type Task struct {
	ID         string
	WorkflowID string
	UserID     string
	RequestID  string
	// EnqueueTime time.Time
}
