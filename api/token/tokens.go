package token

import (
	"net/http"

	"github.com/kaibling/apiforge/apictx"
	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"

	"github.com/kaibling/iggy/initservice"
)

func getTokens(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	ts := initservice.NewTokenService(r.Context(), apictx.GetValue(r.Context(), "user_name").(string))
	tokens, err := ts.ListTokens()
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(tokens).Finish(w, r)
}
