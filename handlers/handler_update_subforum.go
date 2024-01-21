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

// HandlerUpdateSubforum handles the request to update a specific existing subforum.
func HandlerUpdateSubforum(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	subforumTable := os.Getenv("DB_SUBFORUMS_TABLE")
	if subforumTable == "" {
		util.RespondWithError(w, http.StatusInternalServerError, "DB_SUBFORUMS_TABLE is not set in the environment")
		return
	}

	// Retrieve the subforumID from the URL query
	subforumID := chi.URLParam(r, "subforumID")

	// Retrieve the claims from the middleware context (util.AuthenticateUserMiddleware)
	claims, ok := r.Context().Value("userClaims").(*jwt.RegisteredClaims)
	if !ok {
		util.RespondWithError(w, http.StatusInternalServerError, "Error processing user data")
		return
	}

	// // Check if the JWT Subject is a SUPERUSER
	var userType string
	userType, err := util.QueryUserType(claims.Subject)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error querying user's type: \n%v", err))
	}
	if userType != "super" {
		util.RespondWithError(w, http.StatusUnauthorized, "User does not have authority to update this reply")
		return
	}

	// Decode the JSON request body
	type UpdateSubforumRequestData struct {
		NewPhotoURL    string `json:"newPhotoURL"`
		NewName        string `json:"newName"`
		NewDescription string `json:"newDescription"`
	}
	var requestData UpdateSubforumRequestData
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Update name if provided
	if requestData.NewName != "" {
		updateQuery := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", subforumTable)
		_, err := database.GetDB().Exec(updateQuery, requestData.NewName, subforumID)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating subforum name: \n%v", err))
			return
		}
	}

	// Update description if provided
	if requestData.NewDescription != "" {
		updateQuery := fmt.Sprintf("UPDATE %s SET description = $1 WHERE id = $2", subforumTable)
		_, err := database.GetDB().Exec(updateQuery, requestData.NewDescription, subforumID)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating subforum description: \n%v", err))
			return
		}
	}

	// Update photoURL if provided
	if requestData.NewPhotoURL != "" {
		updateQuery := fmt.Sprintf("UPDATE %s SET photo_url = $1 WHERE id = $2", subforumTable)
		_, err := database.GetDB().Exec(updateQuery, requestData.NewPhotoURL, subforumID)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating subforum photo URL: \n%v", err))
			return
		}
	}

	// Respond with a success message
	util.RespondWithJSON(w, http.StatusOK, struct{ Message string }{"Subforum updated successfully"})
}
