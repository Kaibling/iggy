package run

import (
	"net/http"

	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
)

func fetchRun(w http.ResponseWriter, r *http.Request) {
	runID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)

	us, err := bootstrap.NewRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	run, err := us.FetchRun(runID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	e.SetResponse(run).Finish(w, r)
}

func createRun(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)

	var postRun entity.NewRun
	if err := route.ReadPostData(r, &postRun); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	us, err := bootstrap.NewRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	rs, err := bootstrap.NewRunLogService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	err = us.CreateRun(postRun, rs)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	e.SetSuccess().Finish(w, r)
}

func fetchRunLogsByRun(w http.ResponseWriter, r *http.Request) {
	runID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)

	us, err := bootstrap.NewRunLogService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	runLogs, err := us.FetchRunLogsByRun(runID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	e.SetResponse(runLogs).Finish(w, r)
}
