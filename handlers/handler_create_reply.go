package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/util"
)

// HandlerCreateReply handles the request to create a new reply.
func HandlerCreateReply(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	repliesTable := os.Getenv("DB_REPLIES_TABLE")
	if repliesTable == "" {
		util.RespondWithError(w, http.StatusInternalServerError, "DB_REPLIES_TABLE is not set in the environment")
		return
	}

	// Retrieve the claims from the middleware context (util.AuthenticateUserMiddleware)
	claims, ok := r.Context().Value("userClaims").(*jwt.RegisteredClaims)
	if !ok {
		util.RespondWithError(w, http.StatusInternalServerError, "Error processing user data")
		return
	}

	// Decode the JSON request body into CreateThreadRequest struct
	type CreateThreadRequestData struct {
		Content string `json:"content"`
	}
	var requestData CreateThreadRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: \n%v", err))
		return
	}
	defer r.Body.Close()

	// Check for empty title or content
	if requestData.Content == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Content is required")
		return
	}

	// Get threadID from the URL parameter
	threadID := chi.URLParam(r, "threadID")

	// Insert the new thread into the database
	query := fmt.Sprintf("INSERT INTO %s (thread_id, content, created_by, updated_at) VALUES ($1, $2, $3, NOW())", repliesTable)
	_, err = database.GetDB().Exec(query, threadID, requestData.Content, claims.Subject)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"Reply created successfully"})
}
