package adapter

import "github.com/kaibling/iggy/entity"

type ExecutionResult struct {
	Error      error
	SharedData map[string]any
	Logs       []entity.NewRunLog
}
