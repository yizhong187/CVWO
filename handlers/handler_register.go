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

func HandlerRegister(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	usersTable := os.Getenv("DB_TESTING_USERS_TABLE")
	if usersTable == "" {
		log.Fatal("usersTable is not set in the environment")
	}

	// Decode the JSON request body into CreateRequestData struct
	type CreateRequestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var requestData CreateRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check for empty name or description
	if requestData.Username == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Username is required")
		return
	} else if requestData.Password == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Password is required")
	}

	taken, err := util.QueryUsernameTaken(requestData.Username)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	} else if taken {
		util.RespondWithError(w, http.StatusBadRequest, "Username is taken")
		return
	}

	hash, err := util.HashPassword(requestData.Password)

	// // Construct and execute SQL query to insert new user
	query := fmt.Sprintf("INSERT INTO %s (username, type, password_hash) VALUES ($1, $2, $3)", usersTable)
	_, err = database.GetDB().Exec(query, requestData.Username, "normal", hash)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"User created successfully"})
}