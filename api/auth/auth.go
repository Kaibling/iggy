package auth

import (
	"net/http"

	apierror "github.com/kaibling/apiforge/error"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/config"
)

func authLogin(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)

	var postLogin entity.Login
	if err := route.ReadPostData(r, &postLogin); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	us := bootstrap.NewUserServiceAnonym(r.Context(), config.SYSTEM_USER)
	ts := bootstrap.NewTokenServiceAnonym(r.Context(), config.SYSTEM_USER)
	token, err := us.Login(postLogin, ts)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}

	e.SetResponse(token).Finish(w, r)
}

func authLogout(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)

	ts := bootstrap.NewTokenServiceAnonym(r.Context(), config.SYSTEM_USER)
	err := ts.DeleteTokenByValue(ctxkeys.GetValue(r.Context(), ctxkeys.TokenKey).(string))
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}

	e.SetSuccess().Finish(w, r)
}

func authCheck(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)

	ts := bootstrap.NewTokenServiceAnonym(r.Context(), config.SYSTEM_USER)
	// TODO check expiration
	t, err := ts.ReadTokenByValue(ctxkeys.GetValue(r.Context(), ctxkeys.TokenKey).(string))
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}

	e.SetResponse(t).Finish(w, r)
}
