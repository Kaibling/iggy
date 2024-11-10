package adapter

import (
	"time"

	"github.com/dop251/goja"
	"github.com/kaibling/iggy/entity"
)

type Javascript struct {
	logs []entity.NewRunLog
}

func NewJavascriptAdapter() *Javascript {
	return &Javascript{logs: []entity.NewRunLog{}}
}

func (j *Javascript) Execute(code string, sharedData map[string]any) ExecutionResult {
	vm := goja.New()

	vars := map[string]any{
		"shared_data": sharedData,
		"log":         j.log,
	}

	if err := setVariables(vm, vars); err != nil {
		return ExecutionResult{err, sharedData, j.logs}
	}

	_, err := vm.RunString(code)
	if err != nil {
		return ExecutionResult{err, sharedData, j.logs}
	}

	return ExecutionResult{err, sharedData, j.logs}
}

func (j *Javascript) log(msg string) {
	j.logs = append(j.logs, entity.NewRunLog{
		Message:   msg,
		Timestamp: time.Now(),
	})
}

func setVariables(vm *goja.Runtime, variables map[string]any) error {
	for k, v := range variables {
		if err := vm.Set(k, v); err != nil {
			return err
		}
	}
	return nil
}
