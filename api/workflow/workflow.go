package workflow

import (
	"context"
	"net/http"

	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/adapters/broker"
	"github.com/kaibling/iggy/bootstrap"
	bootstrap_broker "github.com/kaibling/iggy/bootstrap/broker"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/utility"
)

func fetchWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e, l := envelope.GetEnvelopeAndLogger(r)

	us, err := bootstrap.NewWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	workflow, err := us.FetchWorkflows([]string{workflowID})
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(workflow).Finish(w, r, l)
}

func createWorkflow(w http.ResponseWriter, r *http.Request) {
	e, l := envelope.GetEnvelopeAndLogger(r)

	var postWorkflow entity.NewWorkflow
	if err := route.ReadPostData(r, &postWorkflow); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if err := postWorkflow.Validate(); err != nil {
		e.SetError(err).Finish(w, r, l)

		return
	}

	postWorkflow.ID = ""
	postWorkflow.BuildIn = false

	us, err := bootstrap.NewWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	newWorkflow, err := us.CreateWorkflow(postWorkflow)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(newWorkflow).Finish(w, r, l)
}

func deleteWorkflow(w http.ResponseWriter, r *http.Request) {
	e, l := envelope.GetEnvelopeAndLogger(r)
	workflowID := route.ReadUrlParam("id", r)

	us, err := bootstrap.NewWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if err := us.DeleteWorkflow(workflowID); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}

func fetchRunsByWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e, l := envelope.GetEnvelopeAndLogger(r)

	us, err := bootstrap.NewRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	runs, err := us.FetchRunByWorkflow(workflowID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(runs).Finish(w, r, l)
}

func executeWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e, l := envelope.GetEnvelopeAndLogger(r)

	logger, err := bootstrap.ContextLogger(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	requestID, err := bootstrap.ContextRequestID(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	cfg, db, username, err := bootstrap.ContextDefaultData(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	subConfig := broker.SubscriberConfig{
		Config:   cfg,
		Username: username,
		DBPool:   db,
	}

	// build adapter
	pub, err := bootstrap_broker.NewPublisher(subConfig, "loopback", logger)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	// TODO check, if workflow exits

	// task to byte
	task, err := utility.EncodeToBytes(entity.Task{
		WorkflowID: workflowID,
		RequestID:  requestID,
	})
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	// send task
	if err := pub.Publish(context.TODO(), "TODO", task); err != nil { //nolint: contextcheck
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}

func patchWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadUrlParam("id", r)
	e, l := envelope.GetEnvelopeAndLogger(r)

	var updateWorkflow entity.UpdateWorkflow
	if err := route.ReadPostData(r, &updateWorkflow); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	ws, err := bootstrap.NewWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	updatedWorkflow, err := ws.Update(workflowID, updateWorkflow)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(updatedWorkflow).Finish(w, r, l)
}
