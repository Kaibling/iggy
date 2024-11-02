package main

import (
	"fmt"

	"github.com/kaibling/iggy/pkg/workflow"
)

func main() {
	// if err := api.Start(); err != nil {
	// 	fmt.Println(err)
	// }

	e := workflow.Engine{}

	wf := workflow.Workflow{Code: "shared_data['number'] = 3; f log('4'); 1 + 1", FailOnError: true, ObjectType: workflow.Javascript}
	wf2 := workflow.Workflow{Code: "log(shared_data['number']);", ObjectType: workflow.Javascript}
	wfolder := workflow.Workflow{ObjectType: workflow.Folder, Children: []workflow.Workflow{wf, wf2}}

	res := e.Execute(wfolder)
	if res.Error != nil {
		fmt.Printf("error: %s\n", res.Error.Error())
	}
	fmt.Printf("good: %s\n", res.Runs)
}
