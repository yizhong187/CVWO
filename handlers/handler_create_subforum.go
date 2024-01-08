package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"

	"github.com/yizhong187/CVWO/util"
)

// HandlerAllUsers handles the request to create a new subforum.
func HandlerCreateSubforum(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	subforumTable := os.Getenv("DB_SUBFORUMS_TABLE")
	if subforumTable == "" {
		log.Fatal("DB_SUBFORUMS_TABLE is not set in the environment")
	}

	// Query the database for the user
	userName := chi.URLParam(r, "name")
	user, err := util.QueryUser(userName)
	if err != nil {
		if err.Error() == "User not found" {
			util.RespondWithError(w, http.StatusNotFound, err.Error())
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Check if user has "superuser" type to create subforum
	if user.Type != "superuser" {
		util.RespondWithError(w, http.StatusForbidden, "User does not have the required permissions")
		return
	}

	// Decode the JSON request body into CreateSubforumRequestData struct
	type CreateSubforumRequestData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	var requestData CreateSubforumRequestData
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Construct and execute SQL query to insert new subforum
	query := fmt.Sprintf("INSERT INTO %s (name, description, created_by, updated_at) VALUES ($1, $2, $3, NOW())", subforumTable)
	_, err = database.GetDB().Exec(query, requestData.Name, requestData.Description, user.ID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"Subforum created successfully"})
}
