package workflow

import (
	"net/http"

	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/route"

	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/service/bootstrap"
)

// func fetchAll(w http.ResponseWriter, r *http.Request) {
// 	e := envelope.ReadEnvelope(r)
// 	us := bootstrap.NewWorkflowService(r.Context())
// 	workflows, err := us.FetchAll()
// 	if err != nil {
// 		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
// 		return
// 	}
// 	e.SetResponse(workflows)
// 	e.Finish(w, r)
// }

func fetchWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)
	us := bootstrap.NewWorkflowService(r.Context())
	workflow, err := us.FetchWorkflows([]string{workflowID})
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(workflow).Finish(w, r)
}

func createWorkflow(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	var postWorkflow entity.NewWorkflow
	if err := route.ReadPostData(r, &postWorkflow); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	if err := postWorkflow.Validate(); err != nil {
		e.SetError(err).Finish(w, r)
		return
	}
	postWorkflow.ID = ""
	postWorkflow.BuildIn = false
	us := bootstrap.NewWorkflowService(r.Context())
	newWorkflow, err := us.CreateWorkflow(postWorkflow)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(newWorkflow).Finish(w, r)
}

func deleteWorkflow(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	workflowID := route.ReadUrlParam("id", r)
	us := bootstrap.NewWorkflowService(r.Context())
	if err := us.DeleteWorkflow(workflowID); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetSuccess().Finish(w, r)
}

func fetchRunsbyWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)
	us := bootstrap.NewRunService(r.Context())
	runs, err := us.FetchRunByWorkflow(workflowID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(runs).Finish(w, r)
}

func executeWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)
	wfs := bootstrap.NewWorkflowService(r.Context())
	rs := bootstrap.NewRunService(r.Context())
	wfes := bootstrap.NewWorkflowEngineService(r.Context())
	_, err := wfs.Execute(workflowID, wfes, rs)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetSuccess().Finish(w, r)
}

func patchWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)
	var updateWorkflow entity.UpdateWorkflow
	if err := route.ReadPostData(r, &updateWorkflow); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	ws := bootstrap.NewWorkflowService(r.Context())
	updatedWorkflow, err := ws.Update(workflowID, updateWorkflow)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(updatedWorkflow).Finish(w, r)
}
