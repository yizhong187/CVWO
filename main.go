package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/routers"
)

func main() {

	database.InitDB()

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	// v1Router := chi.NewRouter()
	// v1Router.Get("/healthz", handlers.HandlerReadiness)
	// v1Router.Get("/err", handlers.HandlerErr)
	// v1Router.Get("/user", handlers.HandlerUser)
	// v1Router.Post("/createuser", handlers.HandlerCreateUser)
	// v1Router.Get("/testing", handlers.HandlerTest)
	// router.Mount("/v1", v1Router)

	v2Router := routers.UserRouter()
	router.Mount("/v2", v2Router)

	// Start the HTTP server on port 8080 and log any errors
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
