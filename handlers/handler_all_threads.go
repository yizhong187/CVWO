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
	usersTable := os.Getenv("DB_USERS_TABLE")
	repliesTable := os.Getenv("DB_REPLIES_TABLE")

	query := fmt.Sprintf(`
		SELECT t.id, t.subforum_id, t.title, t.content, t.created_by, u.name, t.is_pinned, t.created_at, t.updated_at, COUNT(r.id) AS reply_count
		FROM %s t
		INNER JOIN %s u ON t.created_by = u.id
		LEFT JOIN %s r ON t.id = r.thread_id
		WHERE t.subforum_id = $1
		GROUP BY t.id, u.name
		ORDER BY t.id DESC
		`, threadsTable, usersTable, repliesTable)

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
		err := rows.Scan(&thread.ID, &thread.SubforumID, &thread.Title, &thread.Content, &thread.CreatedBy,
			&thread.CreatedByName, &thread.IsPinned, &thread.CreatedAt, &thread.UpdatedAt, &thread.ReplyCount)
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
