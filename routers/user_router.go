package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/yizhong187/CVWO/handlers"
	"github.com/yizhong187/CVWO/util"
)

func UserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/healthz", handlers.HandlerReadiness) // Health check endpoint to verify the service status
		r.Get("/err", handlers.HandlerErr)           // Endpoint to test error handling

		// User-related endpoints for retrieving and updating user information.
		r.With(util.AuthenticateUserMiddleware).Get("/", handlers.HandlerUser)
		r.With(util.AuthenticateUserMiddleware).Put("/", handlers.HandlerUpdateUser)

		// For retrieving all posts (threads and replies) by existing user.
		r.Get("/{userName}/posts", handlers.HandlerUserPosts)

	})

	return r
}
