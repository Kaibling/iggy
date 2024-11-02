package workflow

import (
	"net/http"

	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/route"

	"github.com/kaibling/iggy/model"
	"github.com/kaibling/iggy/service/bootstrap"
)

func fetchAll(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	us := bootstrap.NewWorkflowService(r.Context())
	workflows, err := us.FetchAll()
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(workflows)
	e.Finish(w, r)
}

func fetchWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)
	us := bootstrap.NewWorkflowService(r.Context())
	workflow, err := us.FetchWorkflow(workflowID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(workflow).Finish(w, r)
}

func createWorkflow(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	var postWorkflow model.NewWorkflow
	if err := route.ReadPostData(r, &postWorkflow); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
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
