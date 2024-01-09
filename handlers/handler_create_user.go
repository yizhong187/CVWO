package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"

	"github.com/yizhong187/CVWO/util"
)

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		log.Fatal("usersTable is not set in the environment")
	}

	// Decode the JSON request body into CreateRequestData struct
	type CreateRequestData struct {
		Name string `json:"name"`
	}
	var requestData CreateRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check for empty name or description
	if requestData.Name == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Username is required")
		return
	}

	// // Construct and execute SQL query to insert new user
	query := fmt.Sprintf("INSERT INTO %s (name, type) VALUES ($1, $2)", usersTable)
	_, err = database.GetDB().Exec(query, requestData.Name, "normal")
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"User created successfully"})
}
