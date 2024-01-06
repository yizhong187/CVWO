package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handlerUser(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")

	userDataTable := os.Getenv("DB_USERDATA_TABLE")

	userID := r.URL.Query().Get("id")
	if userID == "" {
		respondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	var user User

	err := db.QueryRow("SELECT id, name FROM $1 WHERE id = $2", userDataTable, userID).Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "User not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	// Respond with the user data in JSON format
	respondWithJSON(w, http.StatusOK, user)
}
