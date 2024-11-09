package adapter

import (
	"fmt"
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
	// TODO error hanndling... really???.....
	err := vm.Set("shared_data", sharedData)
	if err != nil {
		return ExecutionResult{err, sharedData, j.logs}
	}
	err = vm.Set("log", j.log)
	if err != nil {
		return ExecutionResult{err, sharedData, j.logs}
	}
	_, err = vm.RunString(code)
	if err != nil {
		return ExecutionResult{err, sharedData, j.logs}
	}

	return ExecutionResult{err, sharedData, j.logs}
}

func (j *Javascript) log(msg string) {

	fmt.Printf("-> log: %s\n", msg)
	j.logs = append(j.logs, entity.NewRunLog{
		Message:   msg,
		Timestamp: time.Now(),
	})

}
