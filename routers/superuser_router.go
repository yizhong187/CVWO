package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/yizhong187/CVWO/handlers"
	"github.com/yizhong187/CVWO/util"
)

func SuperuserRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/healthz", handlers.HandlerReadiness) // Health check endpoint to verify the service status
		r.Get("/err", handlers.HandlerErr)           // Endpoint to test error handling

		// Superuser-related endpoints for creating and updating subforums.
		r.With(util.AuthenticateUserMiddleware).Post("/subforums", handlers.HandlerCreateSubforum)
		r.With(util.AuthenticateUserMiddleware).Put("/subforums/{subforumID}", handlers.HandlerUpdateSubforum)

		r.Get("/{userName}/posts", handlers.HandlerUserPosts)

		// Superuser-related endpoints for editting user types
		//r.Put("/{name}", handlers.HandlerUpdateUserType)

	})

	return r
}
