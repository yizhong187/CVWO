package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/util"
)

func HandlerUpdateUser(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		log.Fatal("usersTable is not set in the environment")
	}

	// Decode the JSON request body into UpdateRequestData struct
	type UpdateRequestData struct {
		OldName     string `json:"oldName"`
		UpdatedName string `json:"updatedName"`
	}
	var requestData UpdateRequestData
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Query the database to find the UUID of the user with the old name
	var userID string
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", usersTable)
	err := database.GetDB().QueryRow(query, requestData.OldName).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespondWithError(w, http.StatusNotFound, "User not found")
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error querying user: \n%v", err))
		}
		return
	}

	// Execute the update query using the UUID
	updateQuery := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", usersTable)
	_, err = database.GetDB().Exec(updateQuery, requestData.UpdatedName, userID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error updating user: \n%v", err))
		return
	}

	// Respond with a success message
	util.RespondWithJSON(w, http.StatusOK, struct{ Message string }{"User updated successfully"})
}
