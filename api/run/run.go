package run

import (
	"net/http"

	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/route"

	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/service/bootstrap"
)

func fetchRun(w http.ResponseWriter, r *http.Request) {
	runID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)
	us := bootstrap.NewRunService(r.Context())
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
	us := bootstrap.NewRunService(r.Context())
	newRun, err := us.CreateRun(postRun)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(newRun).Finish(w, r)
}

func fetchRunLogsByRun(w http.ResponseWriter, r *http.Request) {
	runID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)
	us := bootstrap.NewRunLogService(r.Context())
	runLogs, err := us.FetchRunLogsByRun(runID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(runLogs).Finish(w, r)
}
