package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

// HandlerAllThreads handles the request to retrieve all threads of a subforum from the database.
func HandlerAllThreads(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	threadsTable := os.Getenv("DB_THREADS_TABLE")
	query := fmt.Sprintf("SELECT * FROM %s WHERE subforum_id = $1", threadsTable)

	subforumID := chi.URLParam(r, "subforumID")

	// Execute sql query and return a rows result set
	rows, err := database.GetDB().Query(query, subforumID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}
	defer rows.Close()

	// Initialize a slice of threads, scan each row into a thread struct and append into the slice
	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		err := rows.Scan(&thread.ID, &thread.SubforumID, &thread.Title, &thread.Content, &thread.CreatedBy, &thread.IsPinned, &thread.CreatedAt, &thread.UpdatedAt)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error scanning row: \n%v", err))
			return
		}
		threads = append(threads, thread)
	}

	if err = rows.Err(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Row error: \n%v", err))
		return
	}

	// Respond with the list of threads in JSON format
	util.RespondWithJSON(w, http.StatusOK, threads)
}
