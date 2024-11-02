package token

import (
	"net/http"

	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"

	"github.com/kaibling/iggy/service/bootstrap"
)

func getTokens(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	ts := bootstrap.NewTokenService(r.Context())
	tokens, err := ts.ListTokens()
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(tokens).Finish(w, r)
}
