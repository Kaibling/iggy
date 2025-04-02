package uiconfig

import (
	"net/http"

	"github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/iggy/bootstrap"
)

func uiconfigGet(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	uiConfigService, err := bootstrap.BuildUIService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}
	e.SetResponse(uiConfigService.GenerateUIConfigs()).Finish(w, r, l)
}
