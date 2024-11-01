package user

import (
	"github.com/go-chi/chi/v5"
	"github.com/kaibling/iggy/middleware"
)

func Route() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Authentication)
		//r.Use(middleware.Authorization)
		r.Post("/", userPost)
		r.Get("/", usersGet)
		r.Get("/{id}", userGet)
		r.Delete("/{id}", userDel)
		r.Get("/{id}/tokens", getUserToken)
	})
	return r
}
