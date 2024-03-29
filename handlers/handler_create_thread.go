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

// HandlerCreateThread handles the request to create a new thread.
func HandlerCreateThread(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	threadsTable := os.Getenv("DB_THREADS_TABLE")
	if threadsTable == "" {
		util.RespondWithError(w, http.StatusInternalServerError, "DB_THREADS_TABLE is not set in the environment")
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
		Title   string `json:"title"`
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
	if requestData.Title == "" || requestData.Content == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Title and content are required")
		return
	}

	// Get subforumID from the URL parameter
	subforumID := chi.URLParam(r, "subforumID")

	// Insert the new thread into the database
	query := fmt.Sprintf("INSERT INTO %s (subforum_id, title, content, created_by, is_pinned, updated_at) VALUES ($1, $2, $3, $4, $5, NOW())", threadsTable)
	_, err = database.GetDB().Exec(query, subforumID, requestData.Title, requestData.Content, claims.Subject, false)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"Thread created successfully"})
}
