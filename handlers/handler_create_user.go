package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"

	//"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")

	userTable := os.Getenv("DB_USERS_TABLE")

	if userTable == "" {
		log.Fatal("userTable is not set in the environment")
	}

	// Check if request is a POST request
	if r.Method != http.MethodPost {
		util.RespondWithError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	type RequestData struct {
		Name string `json:"name"`
	}

	var requestData RequestData

	// Parsing JSON data from POST method
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Error parsing request body")
		return
	}

	// Insert newUser into the database
	query := fmt.Sprintf("INSERT INTO %s (name, type) VALUES ($1, $2)", userTable)
	_, err = database.GetDB().Exec(query, requestData.Name, "normal")
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"User created successfully"})
}
