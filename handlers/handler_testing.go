package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

func HandlerTest(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")

	usersTable := os.Getenv("DB_USERS_TABLE")
	query := fmt.Sprintf("SELECT * FROM %s WHERE name = $1", usersTable)

	// Extracting the name from the URL path
	userName := chi.URLParam(r, "name")
	if userName == "" {
		util.RespondWithError(w, http.StatusBadRequest, "User's name is required")
		return
	}

	var user models.User
	err := database.GetDB().QueryRow(query, userName).Scan(&user.ID, &user.Name, &user.Type, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		}
		return
	}

	// Respond with the user data in JSON format
	util.RespondWithJSON(w, http.StatusOK, user)
}
