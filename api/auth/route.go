package auth

import (
	"github.com/go-chi/chi/v5"
)

func Route() chi.Router { //nolint: ireturn
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Post("/login", authLogin)
		r.Get("/logout", authLogout)
		r.Get("/check", authCheck)

	})

	return r
}
