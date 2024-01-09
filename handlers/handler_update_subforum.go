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

func HandlerUpdateSubforum(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	subforumTable := os.Getenv("DB_SUBFORUMS_TABLE")
	if subforumTable == "" {
		util.RespondWithError(w, http.StatusInternalServerError, "DB_SUBFORUMS_TABLE is not set in the environment")
		return
	}

	// Decode the JSON request body
	type UpdateSubforumRequestData struct {
		NewName        string `json:"newName"`
		NewDescription string `json:"newDescription"`
	}
	var requestData UpdateSubforumRequestData
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Extract the username and subforumID from the URL parameter
	userName := chi.URLParam(r, "name")
	subforumID := chi.URLParam(r, "subforumID")

	// Check if the user is an admin of the subforum
	isAdmin, err := util.IsAdminOf(userName, subforumID)
	log.Printf("%s's admin status: %v", userName, isAdmin)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if !isAdmin {
		util.RespondWithError(w, http.StatusForbidden, "User does not have the required permissions")
		return
	}

	// Execute the update query
	updateQuery := fmt.Sprintf("UPDATE %s SET name = $1, description = $2 WHERE id = $3", subforumTable)
	_, err = database.GetDB().Exec(updateQuery, requestData.NewName, requestData.NewDescription, subforumID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating subforum: \n%v", err))
		return
	}

	// Respond with a success message
	util.RespondWithJSON(w, http.StatusOK, struct{ Message string }{"Subforum updated successfully"})
}
