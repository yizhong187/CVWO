package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

func HandlerUser(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")

	usersTable := os.Getenv("DB_USERS_TABLE")
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = $1", usersTable)

	userName := r.URL.Query().Get("name")
	if userName == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Username is required")
		return
	}

	var user models.User
	err := database.GetDB().QueryRow(query, userName).Scan(&user.Name, &user.UserType)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	// Respond with the user data in JSON format
	util.RespondWithJSON(w, http.StatusOK, user)
}
