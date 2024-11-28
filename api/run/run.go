package run

import (
	"net/http"

	apierror "github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
)

func fetchRun(w http.ResponseWriter, r *http.Request) {
	runID := route.ReadURLParam("id", r)

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

	runs, err := us.FetchRuns([]string{runID})
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if len(runs) < 1 {
		e.SetError(apierror.New(apierror.ErrDataNotFound,
			apierror.ErrDataNotFound.HTTPStatus())).Finish(w, r, l)

		return
	}

	// todo check length
	e.SetResponse(runs[0]).Finish(w, r, l)
}

func createRun(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	var postRun entity.NewRun
	if err := route.ReadPostData(r, &postRun); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	us, err := bootstrap.BuildRouteRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	rs, err := bootstrap.BuildRouteRunLogService(r.Context())
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
	runID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	us, err := bootstrap.BuildRouteRunLogService(r.Context())
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

	rs, err := bootstrap.BuildRouteRunService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	runs, pageData, err := rs.FetchByPagination(params)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	// queriedRequestID, _ := strings.CutPrefix(params.Filter, "request_id=")

	// run, err := us.FetchRunByRequestID(queriedRequestID)
	// if err != nil {
	// 	e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

	// 	return
	// }

	e.SetResponse(runs).SetPagination(pageData).Finish(w, r, l)
}
