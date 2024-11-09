package run

import (
	"github.com/go-chi/chi/v5"
	"github.com/kaibling/iggy/api/middleware"
)

func Route() chi.Router { //nolint: ireturn
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Authentication)
		//r.Use(middleware.Authorization)
		r.Post("/", createRun)
		r.Get("/{id}", fetchRun)
		r.Get("/{id}/logs", fetchRunLogsByRun)
	})

	return r
}
