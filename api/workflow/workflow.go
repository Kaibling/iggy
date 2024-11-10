package workflow

import (
	"net/http"

	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
)

func fetchWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)

	us, err := bootstrap.NewWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

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

	us, err := bootstrap.NewWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

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

	us, err := bootstrap.NewWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	if err := us.DeleteWorkflow(workflowID); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	e.SetSuccess().Finish(w, r)
}

func fetchRunsByWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)

	us, err := bootstrap.NewRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

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

	wfs, err := bootstrap.NewWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	rs, err := bootstrap.NewRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	rls, err := bootstrap.NewRunLogService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	wfes, err := bootstrap.NewWorkflowEngineService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	if err := wfs.Execute(workflowID, wfes, rs, rls); err != nil {
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

	ws, err := bootstrap.NewWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	updatedWorkflow, err := ws.Update(workflowID, updateWorkflow)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	e.SetResponse(updatedWorkflow).Finish(w, r)
}
