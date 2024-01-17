package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

// HandlerSubforum handles the request to retrieve a specific existing subforum.
func HandlerSubforum(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	subforumTable := os.Getenv("DB_SUBFORUMS_TABLE")
	if subforumTable == "" {
		log.Fatal("subforumTable is not set in the environment")
	}

	// Retrieve subforum name from URL query
	subforumID := chi.URLParam(r, "subforumID")
	if subforumID == "" {
		util.RespondWithError(w, http.StatusBadRequest, "Subforum's name is required")
		return
	}

	// Query the database for the subforum
	var subforum models.Subforum
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", subforumTable)
	err := database.GetDB().QueryRow(query, subforumID).Scan(&subforum.ID, &subforum.Name, &subforum.Description, &subforum.CreatedAt, &subforum.UpdatedAt, &subforum.PhotoUrl)
	if err != nil {
		if err == sql.ErrNoRows {
			util.RespondWithError(w, http.StatusNotFound, "Subforum not found")
		} else {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		}
		return
	}

	// Respond with the subforum data in JSON format
	util.RespondWithJSON(w, http.StatusOK, subforum)
}
