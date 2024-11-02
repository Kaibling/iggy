package adapter

type ExecutionResult struct {
	Error      error
	SharedData map[string]any
}
