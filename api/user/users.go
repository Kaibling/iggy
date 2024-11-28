package user

import (
	"net/http"

	apierror "github.com/kaibling/apiforge/apierror"
	"github.com/kaibling/apiforge/envelope"
	"github.com/kaibling/apiforge/route"
	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
)

func usersGet(w http.ResponseWriter, r *http.Request) {
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

	us, err := bootstrap.BuildRouteUserService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	users, pageData, err := us.FetchByPagination(params)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(users).SetPagination(pageData).Finish(w, r, l)
}

func userGet(w http.ResponseWriter, r *http.Request) {
	userID := route.ReadURLParam("id", r)

	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	us, err := bootstrap.BuildRouteUserService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	user, err := us.FetchUser(userID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(user).Finish(w, r, l)
}

func userPost(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	var postUser entity.NewUser
	if err := route.ReadPostData(r, &postUser); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	postUser.ID = ""

	us, err := bootstrap.BuildRouteUserService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	newUser, err := us.CreateUser(postUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(newUser).Finish(w, r, l)
}

func userDel(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	userID := route.ReadURLParam("id", r)

	us, err := bootstrap.BuildRouteUserService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	if err := us.DeleteUser(userID); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetSuccess().Finish(w, r, l)
}

func getUserToken(w http.ResponseWriter, r *http.Request) {
	e, l, merr := envelope.GetEnvelopeAndLogger(r)
	if merr != nil {
		e.SetError(merr).Finish(w, r, l)

		return
	}

	userID := route.ReadURLParam("id", r)

	ts, err := bootstrap.BuildRouteTokenService(r.Context())
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	// TODO check expiration
	tokens, err := ts.ListUserToken(userID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r, l)

		return
	}

	e.SetResponse(tokens).Finish(w, r, l)
}
