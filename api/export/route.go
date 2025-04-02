package export

import (
	"github.com/go-chi/chi/v5"
	"github.com/kaibling/iggy/api/middleware"
)

func Route() chi.Router { //nolint: ireturn
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.Authentication)
		r.Post("/export", dataExport)
		r.Post("/import", dataImport)

	})

	return r
}
