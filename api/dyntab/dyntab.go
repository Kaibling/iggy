package dynTab

import (
	"net/http"

	apierror "github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
)

func fetchDynTab(w http.ResponseWriter, r *http.Request) {
	dynTabID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	dts, err := bootstrap.NewDynTabService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	dynTabs, err := dts.FetchDynamicTables([]string{dynTabID})
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if len(dynTabs) < 1 {
		e.SetError(apierror.New(apierror.ErrDataNotFound,
			apierror.ErrDataNotFound.HTTPStatus())).Finish(w, r, l)

		return
	}

	// todo check length
	e.SetResponse(dynTabs[0]).Finish(w, r, l)
}

func createDynTabs(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	var postDynTab []entity.NewDynamicTable
	if err := route.ReadPostData(r, &postDynTab); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	dts, err := bootstrap.NewDynTabService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	fetchedDynTabs, err := dts.CreateDynamicTables(postDynTab)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(fetchedDynTabs).Finish(w, r, l)
}

func fetchDynTabs(w http.ResponseWriter, r *http.Request) {
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

	us, err := bootstrap.NewDynTabService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	fetchedDynTabs, pageData, err := us.FetchByPagination(params)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(fetchedDynTabs).SetPagination(pageData).Finish(w, r, l)
}

func fetchDynTabVars(w http.ResponseWriter, r *http.Request) {
	dynTabID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	dts, err := bootstrap.NewDynTabService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	fetchedDynTabs, err := dts.FetchVars(dynTabID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(fetchedDynTabs).Finish(w, r, l)
}

func addDynTabVars(w http.ResponseWriter, r *http.Request) {
	dynTabID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	var postDynTabVar []entity.NewDynamicTableVariable
	if err := route.ReadPostData(r, &postDynTabVar); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	dts, err := bootstrap.NewDynTabService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if err := dts.AddVars(dynTabID, postDynTabVar); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}

func deleteDynTabVar(w http.ResponseWriter, r *http.Request) {
	dynTabVarID := route.ReadURLParam("variable_id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	dts, err := bootstrap.NewDynTabService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if err := dts.RemoveVars([]string{dynTabVarID}); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}
