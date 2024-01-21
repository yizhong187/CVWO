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

// HandlerCreateSubforum handles the request to create a new subforum.
func HandlerCreateSubforum(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	subforumTable := os.Getenv("DB_SUBFORUMS_TABLE")
	if subforumTable == "" {
		log.Fatal("DB_SUBFORUMS_TABLE is not set in the environment")
	}

	// Retrieve the claims from the middleware context (util.AuthenticateUserMiddleware)
	claims, ok := r.Context().Value("userClaims").(*jwt.RegisteredClaims)
	if !ok {
		util.RespondWithError(w, http.StatusInternalServerError, "Error processing user data")
		return
	}

	// Check if the JWT Subject is a SUPERUSER
	var userType string
	userType, err := util.QueryUserType(claims.Subject)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error querying user's type: \n%v", err))
	}
	if userType != "super" {
		util.RespondWithError(w, http.StatusUnauthorized, "User does not have authority to view this info")
		return
	}

	// Decode the JSON request body into CreateSubforumRequestData struct
	type CreateSubforumRequestData struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		PhotoURL    string `json:"photoURL"`
	}
	var requestData CreateSubforumRequestData
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check for empty name or description
	if requestData.Name == "" || requestData.Description == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Name and description are required")
		return
	}

	// Construct and execute SQL query to insert new subforum
	query := fmt.Sprintf("INSERT INTO %s (name, description, photo_url) VALUES ($1, $2, $3)", subforumTable)
	_, err = database.GetDB().Exec(query, requestData.Name, requestData.Description, requestData.PhotoURL)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Respond with success message
	util.RespondWithJSON(w, http.StatusCreated, struct{ Message string }{"Subforum created successfully"})
}
