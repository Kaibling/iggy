package dyntab

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
		r.Post("/", createDynTabs)
		r.Get("/", fetchDynTabs)
		r.Get("/{id}", fetchDynTab)

		r.Get("/{id}/fields", fetchDynFields)
		r.Post("/{id}/fields", addDynFields)
		// r.Patch("/{id}/variables/{variable_id}", fetchDynTabVar)   // update variable of id table
		r.Delete("/{id}/fields/{field_id}", deleteDynFields) // delete variable to id table
	})

	return r
}
