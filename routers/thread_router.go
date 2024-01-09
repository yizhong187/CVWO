package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/yizhong187/CVWO/handlers"
)

func ThreadRouter() *chi.Mux {
	r := chi.NewRouter()

	// Thread-related endpoints
	r.Post("/", handlers.HandlerCreateThread) // Endpoint for creating a new thread
	r.Get("/", handlers.HandlerAllThreads)    // Endpoint for retrieving all threads in a subforum
	//r.Get("/{threadID}", handlers.HandlerThread)      // Endpoint for retrieving a specific thread
	//r.Put("/{threadID}", handlers.HandlerUpdateThread) // Endpoint for updating a specific thread
	//r.Delete("/{threadID}", handlers.HandlerDeleteThread) // Endpoint for deleting a specific thread

	// Additional thread-related endpoints can be added here

	return r
}
