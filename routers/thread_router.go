package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/yizhong187/CVWO/handlers"
	"github.com/yizhong187/CVWO/util"
)

func ThreadRouter() *chi.Mux {
	r := chi.NewRouter()

	// Thread-related endpoints
	r.Get("/", handlers.HandlerAllThreads)       // Endpoint for retrieving all replies in a thread
	r.Get("/{threadID}", handlers.HandlerThread) // Endpoint for retrieving a specific thread

	r.With(util.AuthenticateUserMiddleware).Post("/", handlers.HandlerCreateThread)             // Endpoint for creating a new thread
	r.With(util.AuthenticateUserMiddleware).Put("/{threadID}", handlers.HandlerUpdateThread)    // Endpoint for updating a specific thread
	r.With(util.AuthenticateUserMiddleware).Delete("/{threadID}", handlers.HandlerDeleteThread) // Endpoint for deleting a specific thread

	// Mount ReplyRouter under a specific thread
	r.Mount("/{threadID}/replies", ReplyRouter())
	return r
}
