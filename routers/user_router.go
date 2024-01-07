package routers

import (
	"github.com/go-chi/chi/v5"
	//"github.com/go-chi/chi/v5/middleware"
	"github.com/yizhong187/CVWO/handlers"
)

func UserRouter() *chi.Mux {
	r := chi.NewRouter()

	//r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Get("/healthz", handlers.HandlerReadiness)
		r.Get("/err", handlers.HandlerErr)
		r.Post("/", handlers.HandlerCreateUser)
		r.Get("/", handlers.HandlerAllUsers)
		r.Get("/{name}", handlers.HandlerTest)
		r.Put("/{name}", handlers.HandlerUpdateUser)
		r.Delete("/{name}", handlers.HandlerDeleteUser)
	})

	return r
}
