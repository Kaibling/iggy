package workflow

import (
	"github.com/go-chi/chi/v5"
	"github.com/kaibling/iggy/api/middleware"
)

func Route() chi.Router { //nolint: ireturn
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Authentication)
		// r.Use(middleware.Authorization)
		r.Post("/", createWorkflow)
		r.Patch("/{id}", patchWorkflow)
		r.Get("/{id}", fetchWorkflow)
		r.Delete("/{id}", deleteWorkflow)
		r.Get("/{id}/runs", fetchRunsByWorkflow)
		r.Post("/{id}/execute", executeWorkflow)
	})

	return r
}
