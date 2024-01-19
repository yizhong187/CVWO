package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/util"
)

// HandlerDeleteThread handles the request to delete an existing thread. All replies under the thread will be deleled as well.
func HandlerDeleteThread(w http.ResponseWriter, r *http.Request) {
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

	var userType string
	userType, err = util.QueryUserType(claims.Subject)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error querying user's type: \n%v", err))
	}

	// Check if the JWT Subject matches the original poster or is a SUPERUSER
	if originalPoster != claims.Subject && userType != "super" {
		util.RespondWithError(w, http.StatusUnauthorized, "User does not have authority to delete this thread")
		return
	}

	// Delete the existing thread from the database
	query = fmt.Sprintf("DELETE FROM %s WHERE id = $1", threadsTable)
	_, err = database.GetDB().Exec(query, threadID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusOK, struct{ Message string }{"Thread deleted successfully"})
}
