package run

import (
	"github.com/go-chi/chi/v5"
	apimiddleware "github.com/kaibling/apiforge/middleware"
	"github.com/kaibling/iggy/api/middleware"
)

func Route() chi.Router { //nolint:nolintlint,ireturn
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Authentication)
		// r.Use(middleware.Authorization)
		r.Use(apimiddleware.ParsePagination)
		r.Post("/", createRun)
		r.Get("/", fetchRuns)
		r.Get("/{id}", fetchRun)
		r.Get("/{id}/logs", fetchRunLogsByRun)
	})

	return r
}
