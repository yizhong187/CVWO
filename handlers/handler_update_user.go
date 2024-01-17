package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/util"
)

// HandlerUpdateUser handles the request to update the details of an existing user.
func HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		log.Fatal("usersTable is not set in the environment")
	}

	// Retrieve the claims from the middleware context (util.AuthenticateUserMiddleware)
	claims, ok := r.Context().Value("userClaims").(*jwt.RegisteredClaims)
	if !ok {
		util.RespondWithError(w, http.StatusInternalServerError, "Error processing user data")
		return
	}

	// Decode the JSON request body into UpdateRequestData struct
	type UpdateRequestData struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	var requestData UpdateRequestData
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Hash password using util.HashPassword
	hash, err := util.HashPassword(requestData.Password)

	// Construct and execute SQL query to update the user
	updateQuery := fmt.Sprintf("UPDATE %s SET name = $1, password_hash = $2 WHERE id = $3", usersTable)
	_, err = database.GetDB().Exec(updateQuery, requestData.Name, hash, claims.Subject)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating user: \n%v", err))
		return
	}

	// Respond with a success message
	util.RespondWithJSON(w, http.StatusOK, struct{ Message string }{"User updated successfully"})
}
