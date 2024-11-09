package user

import (
	"net/http"

	"github.com/kaibling/apiforge/ctxkeys"
	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/params"
	"github.com/kaibling/apiforge/route"

	"github.com/kaibling/iggy/bootstrap"
	"github.com/kaibling/iggy/entity"
)

func usersGet(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	params := ctxkeys.GetValue(r.Context(), ctxkeys.PaginationKey).(params.Pagination)
	us := bootstrap.NewUserService(r.Context())
	users, pageData, err := us.FetchByPagination(params)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(users).SetPagination(pageData).Finish(w, r)
}

func userGet(w http.ResponseWriter, r *http.Request) {
	userID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)
	us := bootstrap.NewUserService(r.Context())
	user, err := us.FetchUser(userID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(user).Finish(w, r)
}

func userPost(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	var postUser entity.NewUser
	if err := route.ReadPostData(r, &postUser); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	postUser.ID = ""
	us := bootstrap.NewUserService(r.Context())
	newUser, err := us.CreateUser(postUser)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(newUser).Finish(w, r)
}

func userDel(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	userID := route.ReadUrlParam("id", r)
	us := bootstrap.NewUserService(r.Context())
	if err := us.DeleteUser(userID); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetSuccess().Finish(w, r)
}

func getUserToken(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	userID := route.ReadUrlParam("id", r)
	ts := bootstrap.NewTokenService(r.Context())
	// TODO check expiration
	tokens, err := ts.ListUserToken(userID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(tokens).Finish(w, r)
}
