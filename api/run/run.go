package run

import (
	"net/http"
	"strings"

	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
)

func fetchRun(w http.ResponseWriter, r *http.Request) {
	runID := route.ReadUrlParam("id", r)
	e, l := envelope.GetEnvelopeAndLogger(r)

	us, err := bootstrap.NewRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	run, err := us.FetchRun(runID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(run).Finish(w, r, l)
}

func createRun(w http.ResponseWriter, r *http.Request) {
	e, l := envelope.GetEnvelopeAndLogger(r)

	var postRun entity.NewRun
	if err := route.ReadPostData(r, &postRun); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	us, err := bootstrap.NewRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	rs, err := bootstrap.NewRunLogService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	err = us.CreateRun(postRun, rs)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}

func fetchRunLogsByRun(w http.ResponseWriter, r *http.Request) {
	runID := route.ReadUrlParam("id", r)
	e, l := envelope.GetEnvelopeAndLogger(r)

	us, err := bootstrap.NewRunLogService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	runLogs, err := us.FetchRunLogsByRun(runID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(runLogs).Finish(w, r, l)
}

func fetchRuns(w http.ResponseWriter, r *http.Request) {
	e, l := envelope.GetEnvelopeAndLogger(r)

	params, err := bootstrap.ContextParams(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	us, err := bootstrap.NewRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	queriedRequestID, _ := strings.CutPrefix(params.Filter, "request_id=")

	run, err := us.FetchRunByRequestID(queriedRequestID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(run).Finish(w, r, l)
}
