package export

import (
	"net/http"

	apierror "github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/iggy/apperror"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/pkg/config"
)

func dataExport(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	// var postBackup entity.Backup
	// if err := route.ReadPostData(r, &postBackup); err != nil {
	// 	e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

	// 	return
	// }
	cfg, ok := ctxkeys.GetValue(r.Context(), ctxkeys.AppConfigKey).(config.Configuration)
	if !ok {
		e.SetError(apperror.ErrNewMissingContext("config")).Finish(w, r, l)

		return
	}

	wfs, err := bootstrap.BuildRouteWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	err = wfs.ExportToDir(cfg.App.ExportLocalPath)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)

}

func dataImport(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	cfg, ok := ctxkeys.GetValue(r.Context(), ctxkeys.AppConfigKey).(config.Configuration)
	if !ok {
		e.SetError(apperror.ErrNewMissingContext("config")).Finish(w, r, l)

		return
	}

	wfs, err := bootstrap.BuildRouteWorkflowService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	err = wfs.ImportFromFiles(cfg.App.ImportLocalPath)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)

}
