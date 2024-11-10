package auth

import (
	"net/http"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
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

	us, err := bootstrap.NewUserServiceAnonym(r.Context(), config.SystemUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	ts, err := bootstrap.NewTokenServiceAnonym(r.Context(), config.SystemUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	token, err := us.Login(postLogin, ts)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	e.SetResponse(token).Finish(w, r)
}

func authLogout(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)

	ts, err := bootstrap.NewTokenServiceAnonym(r.Context(), config.SystemUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	err = ts.DeleteTokenByValue(ctxkeys.GetValue(r.Context(), ctxkeys.TokenKey).(string))
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	e.SetSuccess().Finish(w, r)
}

func authCheck(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)

	ts, err := bootstrap.NewTokenServiceAnonym(r.Context(), config.SystemUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	// TODO check expiration
	t, err := ts.ReadTokenByValue(ctxkeys.GetValue(r.Context(), ctxkeys.TokenKey).(string))
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)

		return
	}

	e.SetResponse(t).Finish(w, r)
}
