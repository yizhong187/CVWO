package routers

import (
	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	"github.com/yizhong187/CVWO/handlers"
)

func SubforumRouter() *chi.Mux {
	r := chi.NewRouter()

	//r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)

	// Subforum-related endpoints for creating, retrieving, updating, and deleting subforums
	r.Post("/", handlers.HandlerCreateSubforum)
	r.Get("/", handlers.HandlerAllSubforums)
	r.Get("/{subforumID}", handlers.HandlerSubforum)
	r.Put("/{subforumID}", handlers.HandlerTesting)
	// r.Delete("/{subforumID}", handlers.HandlerDeleteSubforum)

	return r
}
