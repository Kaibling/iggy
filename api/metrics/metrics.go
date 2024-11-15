package metrics

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Route() chi.Router { //nolint: ireturn,nolintlint
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Handle("/", promhttp.Handler())
	})

	return r
}
