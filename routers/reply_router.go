package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/yizhong187/CVWO/handlers"
	"github.com/yizhong187/CVWO/util"
)

func ReplyRouter() *chi.Mux {
	r := chi.NewRouter()

	// Thread-related endpoints
	r.Get("/", handlers.HandlerAllReplies)                                                    // Endpoint for retrieving all replies in a thread
	r.With(util.AuthenticateUserMiddleware).Post("/", handlers.HandlerCreateReply)            // Endpoint for creating a new reply
	r.With(util.AuthenticateUserMiddleware).Put("/{replyID}", handlers.HandlerUpdateReply)    // Endpoint for updating a specific reply
	r.With(util.AuthenticateUserMiddleware).Delete("/{replyID}", handlers.HandlerDeleteReply) // Endpoint for deleting a specific reply

	return r
}
