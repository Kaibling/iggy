package user

import (
	"github.com/go-chi/chi/v5"
	apimiddleware "github.com/kaibling/apiforge/middleware"
	"github.com/kaibling/iggy/api/middleware"
)

func Route() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Authentication)
		//r.Use(middleware.Authorization)

		r.Use(apimiddleware.ParsePagination)
		r.Post("/", userPost)
		r.Get("/", usersGet)
		r.Get("/{id}", userGet)
		r.Delete("/{id}", userDel)
		r.Get("/{id}/tokens", getUserToken)
	})
	return r
}
