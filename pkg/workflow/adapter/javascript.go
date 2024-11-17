package adapter

import (
	"fmt"
	"time"

	"github.com/dop251/goja"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/utility"
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
		"log_obj":     j.logObject,
		"fetch":       utility.JSONFetch,
		// "save": ,
	}

	if err := setVariables(vm, vars); err != nil {
		return ExecutionResult{err, sharedData, j.logs}
	}

	result, err := vm.RunString(code)
	if err != nil {
		return ExecutionResult{err, sharedData, j.logs}
	}

	fmt.Println(result.Export()) //nolint: forbidigo

	return ExecutionResult{err, sharedData, j.logs}
}

func (j *Javascript) log(msg string) {
	j.logs = append(j.logs, entity.NewRunLog{
		Message:   msg,
		Timestamp: time.Now(),
	})
}

func (j *Javascript) logObject(obj any) {
	j.logs = append(j.logs, entity.NewRunLog{
		Message:   utility.Pretty(obj),
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
