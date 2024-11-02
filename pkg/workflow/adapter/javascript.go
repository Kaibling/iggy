package adapter

import (
	"fmt"

	"github.com/dop251/goja"
)

type Javascript struct {
}

func NewJavascriptAdapter() Javascript {
	return Javascript{}
}
func (j Javascript) Execute(code string, sharedData map[string]any) ExecutionResult {
	vm := goja.New()

	vm.Set("shared_data", sharedData)
	vm.Set("log", log)
	_, err := vm.RunString(code)
	if err != nil {
		return ExecutionResult{err, sharedData}
	}

	return ExecutionResult{err, sharedData}
}

func log(s string) {
	fmt.Printf("-> log: %s\n", s)
}
