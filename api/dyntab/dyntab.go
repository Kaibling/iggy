package dyntab

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

	dts, err := bootstrap.BuildRouteDynTabService(r.Context())
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

	dts, err := bootstrap.BuildRouteDynTabService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	dss, err := bootstrap.BuildRouteDynSchemaService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	fetchedDynTabs, err := dts.CreateDynamicTables(postDynTab, dss)
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

	us, err := bootstrap.BuildRouteDynTabService(r.Context())
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

func fetchDynFields(w http.ResponseWriter, r *http.Request) {
	dynTabID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	dts, err := bootstrap.BuildRouteDynFieldService(r.Context())
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

func addDynFields(w http.ResponseWriter, r *http.Request) {
	// todo improve api response
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	var postDynTabVar []entity.NewDynamicField
	if err := route.ReadPostData(r, &postDynTabVar); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	dfs, err := bootstrap.BuildRouteDynFieldService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if err := dfs.AddFields(postDynTabVar); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}

func deleteDynFields(w http.ResponseWriter, r *http.Request) {
	dynTabID := route.ReadURLParam("id", r)
	dynFieldID := route.ReadURLParam("field_id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	dfs, err := bootstrap.BuildRouteDynFieldService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	dts, err := bootstrap.BuildRouteDynTabService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	tab, err := dts.FetchDynamicTables([]string{dynTabID})
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if err := dfs.RemoveVars(tab[0].Name, []string{dynFieldID}); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}
