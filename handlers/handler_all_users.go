package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

// HandlerAllUsers handles the request to retrieve all users from the database.
func HandlerAllUsers(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")

	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		util.RespondWithError(w, http.StatusInternalServerError, "DB_USERS_TABLE is not set in the environment")
		return
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

	query := fmt.Sprintf("SELECT id, name, type, created_at FROM %s", usersTable)

	// Execute sql query and return a rows result set
	rows, err := database.GetDB().Query(query)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}
	defer rows.Close()

	// Initialise a slice of users, scan each row into a user struct and append into the slice
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Type, &user.CreatedAt)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error scanning row: \n%v", err))
			return
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Row error: \n%v", err))
		return
	}

	// Respond with list of users in JSON format
	util.RespondWithJSON(w, http.StatusOK, users)
}
