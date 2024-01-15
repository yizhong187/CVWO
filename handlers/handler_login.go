package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/util"
)

func HandlerLogin(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	usersTable := os.Getenv("DB_TESTING_USERS_TABLE")
	if usersTable == "" {
		log.Fatal("usersTable is not set in the environment")
	}
	secretKey := os.Getenv("SECRET_KEY")
	if usersTable == "" {
		log.Fatal("SECRET_KEY is not set in the environment")
	}

	// Decode the JSON request body into CreateRequestData struct
	type CreateRequestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var requestData CreateRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Check for empty name or description
	if requestData.Username == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Username is required")
		return
	} else if requestData.Password == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Password is required")
		return
	}

	// Construct and execute SQL query to retrieve passwordHash
	var passwordHash string
	var userID string
	query := fmt.Sprintf("SELECT password_hash, id FROM %s WHERE username = $1;	", usersTable)
	err = database.GetDB().QueryRow(query, requestData.Username).Scan(&passwordHash, &userID)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespondWithError(w, http.StatusBadRequest, "User not found")
			return
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving passwordHash: %v", err))
			return
		}
	}

	if !util.CheckPasswordHash(requestData.Password, passwordHash) {
		util.RespondWithError(w, http.StatusBadRequest, "Wrong password")
		return
	}

	// Define the standard claims
	claims := &jwt.RegisteredClaims{
		Issuer:    "github.com/yizhong187/CVWO",
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 1 day
	}

	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Could not login")
		http.Error(w, "could not login", http.StatusInternalServerError)
		return
	}

	// Set the token in an HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/", // Make sure the cookie is sent with every request to the server
	})

	util.RespondWithJSON(w, http.StatusOK, "Successfully logged in")

}
