package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

func HandlerTesting(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	usersTable := os.Getenv("DB_TESTING_USERS_TABLE")
	if usersTable == "" {
		log.Fatal("usersTable is not set in the environment")
	}
	secretKey := os.Getenv("SECRET_KEY")
	if usersTable == "" {
		log.Fatal("SECRET_KEY is not set in the environment")
	}

	cookie, err := r.Cookie("jwt")
	if err != nil {
		util.RespondWithError(w, http.StatusUnauthorized, "User unauthenticated")
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			util.RespondWithError(w, http.StatusUnauthorized, "User unauthenticated")
			return
		}
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Bad request: \n%v", err))
		return
	}

	if !token.Valid {
		util.RespondWithError(w, http.StatusUnauthorized, "User unauthenticated")
		return
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		util.RespondWithError(w, http.StatusUnauthorized, "User unauthenticated")
		return
	}

	var user models.TestingUser
	query := fmt.Sprintf("SELECT id, username, type, created_at FROM %s WHERE id = $1", usersTable)
	err = database.GetDB().QueryRow(query, claims.Subject).Scan(&user.ID, &user.Name, &user.Type, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespondWithError(w, http.StatusBadRequest, "User not found")
			return
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving user type: %v", err))
			return
		}
	}

	util.RespondWithJSON(w, http.StatusOK, user)
}

func HandlerLogout(w http.ResponseWriter, r *http.Request) {
	// Set the token in an HTTP-only cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/", // Make sure the cookie is sent with every request to the server
	})

	util.RespondWithJSON(w, http.StatusOK, "Successfully logged out")
}
