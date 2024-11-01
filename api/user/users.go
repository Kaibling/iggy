package user

import (
	"net/http"

	"github.com/kaibling/apiforge/apictx"
	"github.com/kaibling/apiforge/envelope"
	apierror "github.com/kaibling/apiforge/error"
	"github.com/kaibling/apiforge/route"

	"github.com/kaibling/iggy/initservice"
	"github.com/kaibling/iggy/model"
)

func usersGet(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	us := initservice.NewUserService(r.Context(), apictx.GetValue(r.Context(), "user_name").(string))
	users, err := us.FetchAll()
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(users)
	e.Finish(w, r)
}

func userGet(w http.ResponseWriter, r *http.Request) {
	userID := route.ReadUrlParam("id", r)
	e := envelope.ReadEnvelope(r)
	us := initservice.NewUserService(r.Context(), apictx.GetValue(r.Context(), "user_name").(string))
	user, err := us.FetchUser(userID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(user).Finish(w, r)
}

func userPost(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	var postUser model.NewUser
	if err := route.ReadPostData(r, &postUser); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	us := initservice.NewUserService(r.Context(), apictx.GetValue(r.Context(), "user_name").(string))
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
	us := initservice.NewUserService(r.Context(), apictx.GetValue(r.Context(), "user_name").(string))
	if err := us.DeleteUser(userID); err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetSuccess().Finish(w, r)
}

func getUserToken(w http.ResponseWriter, r *http.Request) {
	e := envelope.ReadEnvelope(r)
	userID := route.ReadUrlParam("id", r)
	ts := initservice.NewTokenService(r.Context(), apictx.GetValue(r.Context(), "user_name").(string))
	// TODO check expiration
	tokens, err := ts.ListUserToken(userID)
	if err != nil {
		e.SetError(apierror.NewGeneric(err)).Finish(w, r)
		return
	}
	e.SetResponse(tokens).Finish(w, r)
}
