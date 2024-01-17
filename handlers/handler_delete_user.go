//[DEPRECATED]

package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/util"
)

// HandlerDeleteReply handles the request to delete an existing reply.
func HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		log.Fatal("usersTable is not set in the environment")
	}

	// Retrieve user name from URL path parameter
	name := chi.URLParam(r, "name")
	if name == "" {
		util.RespondWithError(w, http.StatusBadRequest, "User's name is required")
		return
	}

	// Delete the user from the database
	query := fmt.Sprintf("DELETE FROM %s WHERE name = $1", usersTable)
	result, err := database.GetDB().Exec(query, name)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}

	// Check if a row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error checking rows affected: \n%v", err))
		return
	}
	if rowsAffected == 0 {
		util.RespondWithError(w, http.StatusNotFound, "User not found or already deleted")
		return
	}

	// Respond with a success message
	util.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User successfully deleted"})
}
