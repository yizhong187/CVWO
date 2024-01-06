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

	userDataTable := os.Getenv("DB_USERDATA_TABLE")
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id = $1", userDataTable)

	userID := r.URL.Query().Get("id")
	if userID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	var user models.User
	err := database.GetDB().QueryRow(query, userID).Scan(&user.ID, &user.Name)
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
