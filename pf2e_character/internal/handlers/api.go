package handlers

import (
	"pf2e-character-api/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// Middleware is a function that gets called before primary function which handles the endpoint
func Handler(r *chi.Mux) {
	// GLOBAL MIDDLEWARE
	// Removes trailing slash on route which would cause 404 error
	r.Use(chimiddle.StripSlashes)

	r.Route("/v1/token", func(router chi.Router) {
		router.Post("/", CreateUser)
		router.Get("/", GetAuthToken)
	})

	r.Route("/v1/character", func(router chi.Router) {
		router.Use(middleware.Authorization)

		router.Get("/", GetCharacter)
		router.Post("/", CreateCharacter)
		router.Put("/", UpdateCharacter)
		router.Delete("/", DeleteCharacter)
	})

}
