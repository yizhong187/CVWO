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

// HandlerUpdateThread handles the request to update a specific existing thread.
func HandlerUpdateThread(w http.ResponseWriter, r *http.Request) {

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

	// Get threadID from the URL parameter
	threadID := chi.URLParam(r, "threadID")
	if threadID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Thread ID is required")
		return
	}

	// Query the original poster from the threads table
	var originalPoster string
	query := fmt.Sprintf("SELECT created_by FROM %s WHERE id = $1", threadsTable)
	err := database.GetDB().QueryRow(query, threadID).Scan(&originalPoster)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error querying original poster: \n%v", err))
		return
	}

	// Check if the JWT Subject matches the original poster
	if originalPoster != claims.Subject {
		util.RespondWithError(w, http.StatusUnauthorized, "User does not have authority to update this thread")
		return
	}

	// Decode the JSON request body into UpdateThreadRequestData struct
	type UpdateThreadRequestData struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var requestData UpdateThreadRequestData
	err = json.NewDecoder(r.Body).Decode(&requestData)
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

	// Update the existing thread in the database
	query = fmt.Sprintf("UPDATE %s SET title = $2, content = $3, updated_at = NOW() WHERE id = $1 AND created_by = $4", threadsTable)
	_, err = database.GetDB().Exec(query, threadID, requestData.Title, requestData.Content, claims.Subject)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusOK, struct{ Message string }{"Thread updated successfully"})
}
