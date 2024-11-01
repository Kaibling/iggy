package auth

import (
	"net/http"

	apierror "github.com/kaibling/apiforge/error"

	"github.com/kaibling/apiforge/apictx"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/config"
	"github.com/kaibling/iggy/initservice"
	"github.com/kaibling/iggy/model"
)

func authLogin(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	var postLogin model.Login
	if err := route.ReadPostData(r, &postLogin); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	us := initservice.NewUserService(r.Context(), config.SYSTEM_USER)
	ts := initservice.NewTokenService(r.Context(), config.SYSTEM_USER)
	token, err := us.Login(postLogin, ts)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(token).Finish(w, r)
}

func authLogout(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)

	ts := initservice.NewTokenService(r.Context(), config.SYSTEM_USER)
	err := ts.DeleteTokenByValue(apictx.GetValue(r.Context(), "user_token").(string))
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetSuccess().Finish(w, r)
}

func authCheck(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)

	ts := initservice.NewTokenService(r.Context(), config.SYSTEM_USER)
	// TODO check exporation
	t, err := ts.ReadTokenByValue(apictx.GetValue(r.Context(), "user_token").(string))
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(t).Finish(w, r)
}
