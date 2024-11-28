package token

import (
	"net/http"

	apierror "github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/iggy/bootstrap"
)

func getTokens(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	ts, err := bootstrap.BuildRouteTokenService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	tokens, err := ts.ListTokens()
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(tokens).Finish(w, r, l)
}
