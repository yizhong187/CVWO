package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/util"
)

func HandlerUseridToUsername(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	repliesTable := os.Getenv("DB_USERS_TABLE")
	if repliesTable == "" {
		util.RespondWithError(w, http.StatusInternalServerError, "DB_USERS_TABLE is not set in the environment")
		return
	}

	// Decode the JSON request body into CreateUsernameRequest struct
	type CreateUsernameRequestData struct {
		ID string `json:"id"`
	}
	var requestData CreateUsernameRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check for empty content
	if requestData.ID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "ID is required")
		return
	}

	username, err := util.QueryUsername(requestData.ID)
	if err != nil {
		if err.Error() == "User not found!" {
			util.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Respond with the user data in JSON format
	util.RespondWithJSON(w, http.StatusOK, username)
}
