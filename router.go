package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		//r.Post("/", CreateUser)
		//r.Get("/", GetAllUsers)
		//r.Get("/{id}", GetUser)
		//r.Put("/{id}", UpdateUser)
		//r.Delete("/{id}", DeleteUser)
	})

	return r
}
