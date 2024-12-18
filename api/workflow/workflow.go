package workflow

import (
	"context"
	"net/http"

	apierror "github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/bootstrap"
	bootstrap_broker "github.com/kaibling/iggy/bootstrap/broker"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/utility"
	"github.com/kaibling/iggy/service"
)

func fetchWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	us, err := bootstrap.BuildRouteWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	workflows, err := us.FetchWorkflows([]string{workflowID})
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if len(workflows) < 1 {
		e.SetError(apierror.New(apierror.ErrDataNotFound,
			apierror.ErrDataNotFound.HTTPStatus())).Finish(w, r, l)

		return
	}

	e.SetResponse(workflows[0]).Finish(w, r, l)
}

func createWorkflows(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	var postWorkflows []*entity.NewWorkflow
	if err := route.ReadPostData(r, &postWorkflows); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	// todo validate should be business logic, not router logic
	for _, wf := range postWorkflows {
		if err := wf.Validate(); err != nil {
			e.SetError(err).Finish(w, r, l)

			return
		}

		wf.ID = ""
		wf.BuildIn = false
	}

	us, err := bootstrap.BuildRouteWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	newWorkflows, err := us.CreateWorkflows(postWorkflows)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(newWorkflows).Finish(w, r, l)
}

func deleteWorkflow(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	workflowID := route.ReadURLParam("id", r)

	us, err := bootstrap.BuildRouteWorkflowService(r.Context())
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
	workflowID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	us, err := bootstrap.BuildRouteRunService(r.Context())
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

func executeWorkflow(w http.ResponseWriter, r *http.Request) { //nolint: funlen
	workflowID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	errs := apierror.NewMultiError()

	logger, err := bootstrap.ContextLogger(r.Context())
	errs.Add(err)

	requestID, err := bootstrap.ContextRequestID(r.Context())
	errs.Add(err)

	userID, err := bootstrap.ContextUserID(r.Context())
	errs.Add(err)

	cfg, db, _, username, err := bootstrap.ContextDefaultData(r.Context())
	errs.Add(err)

	if errs.HasError() {
		e.SetError(apierror.NewMulti(apierror.ErrContextMissing, errs.GetStrErrors())).Finish(w, r, l)

		return
	}

	brokerConfig := service.Config{
		Config:   cfg,
		Username: username,
		DBPool:   db,
	}

	// build adapter
	pub, err := bootstrap_broker.NewPublisher(brokerConfig, logger)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	// TODO check, if workflow exits

	// task to byte
	task, err := utility.EncodeToBytes(entity.Task{
		WorkflowID: workflowID,
		RequestID:  requestID,
		UserID:     userID,
	})
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	// send task
	if err := pub.Publish(context.TODO(), cfg.Broker.Channel, task); err != nil { //nolint: contextcheck
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}

func patchWorkflow(w http.ResponseWriter, r *http.Request) {
	workflowID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	var updateWorkflow entity.UpdateWorkflow
	if err := route.ReadPostData(r, &updateWorkflow); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	ws, err := bootstrap.BuildRouteWorkflowService(r.Context())
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

func fetchWorkflows(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	params, err := bootstrap.ContextParams(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	us, err := bootstrap.BuildRouteWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	workflows, pageData, err := us.FetchByPagination(params)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(workflows).SetPagination(pageData).Finish(w, r, l)
}
