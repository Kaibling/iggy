package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/kaibling/iggy/api/auth"
	"github.com/kaibling/iggy/api/token"
	"github.com/kaibling/iggy/api/user"
)

func ApiRoute() chi.Router {
	r := chi.NewRouter()
	r.Mount("/auth", auth.Route())
	r.Mount("/users", user.Route())
	r.Mount("/tokens", token.Route())
	return r
}
