package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/util"
)

// HandlerPostReply handles the request to post a new reply to a thread.
func HandlerPostReply(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	repliesTable := os.Getenv("DB_REPLIES_TABLE")
	if repliesTable == "" {
		util.RespondWithError(w, http.StatusInternalServerError, "DB_REPLIES_TABLE is not set in the environment")
		return
	}

	// Decode the JSON request body into CreateReplyRequest struct
	type CreateReplyRequestData struct {
		Content string `json:"content"`
	}
	var requestData CreateReplyRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check for empty content
	if requestData.Content == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Content is required")
		return
	}

	// Get threadID and name (used to find userID) from the URL parameter
	threadID := chi.URLParam(r, "threadID")
	name := chi.URLParam(r, "name")
	userID, err := util.QueryUserID(name)
	if err != nil {
		if err.Error() == "User not found" {
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		}
		return
	}

	// Insert the new reply into the database
	query := fmt.Sprintf("INSERT INTO %s (thread_id, content, created_by, updated_at) VALUES ($1, $2, $3, NOW())", repliesTable)
	_, err = database.GetDB().Exec(query, threadID, requestData.Content, userID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"Reply posted successfully"})
}
