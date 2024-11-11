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
	e, l := envelope.GetEnvelopeAndLogger(r)

	var postLogin entity.Login
	if err := route.ReadPostData(r, &postLogin); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	us, err := bootstrap.NewUserServiceAnonym(r.Context(), config.SystemUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	ts, err := bootstrap.NewTokenServiceAnonym(r.Context(), config.SystemUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	token, err := us.Login(postLogin, ts)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(token).Finish(w, r, l)
}

func authLogout(w http.ResponseWriter, r *http.Request) {
	e, l := envelope.GetEnvelopeAndLogger(r)

	token, err := bootstrap.ContextToken(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	ts, err := bootstrap.NewTokenServiceAnonym(r.Context(), config.SystemUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	err = ts.DeleteTokenByValue(token)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}

func authCheck(w http.ResponseWriter, r *http.Request) {
	e, l := envelope.GetEnvelopeAndLogger(r)

	ts, err := bootstrap.NewTokenServiceAnonym(r.Context(), config.SystemUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	// TODO check expiration
	t, err := ts.ReadTokenByValue(ctxkeys.GetValue(r.Context(), ctxkeys.TokenKey).(string))
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(t).Finish(w, r, l)
}
