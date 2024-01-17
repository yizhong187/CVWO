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

// HandlerDeleteReply handles the request to delete an existing reply.
func HandlerDeleteReply(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")
	repliesTable := os.Getenv("DB_REPLIES_TABLE")
	if repliesTable == "" {
		util.RespondWithError(w, http.StatusInternalServerError, "DB_REPLIES_TABLE is not set in the environment")
		return
	}

	// Retrieve the claims from the middleware context
	claims, ok := r.Context().Value("userClaims").(*jwt.RegisteredClaims)
	if !ok {
		util.RespondWithError(w, http.StatusInternalServerError, "Error processing user data")
		return
	}

	// Get replyID from the URL parameter
	replyID := chi.URLParam(r, "replyID")
	if replyID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Reply ID is required")
		return
	}

	// Query the original poster from the replies table
	var originalPoster string
	query := fmt.Sprintf("SELECT created_by FROM %s WHERE id = $1", repliesTable)
	err := database.GetDB().QueryRow(query, replyID).Scan(&originalPoster)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error querying original poster of the reply: \n%v", err))
		return
	}

	// Check if the JWT Subject matches the original poster
	if originalPoster != claims.Subject {
		util.RespondWithError(w, http.StatusUnauthorized, "User does not have authority to delete this reply")
		return
	}

	// Delete the existing reply from the database
	query = fmt.Sprintf("DELETE FROM %s WHERE id = $1", repliesTable)
	_, err = database.GetDB().Exec(query, replyID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusOK, struct{ Message string }{"Reply deleted successfully"})
}
