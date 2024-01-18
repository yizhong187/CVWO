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

// HandlerThread handles the request to retrieve a specific existing thread.
func HandlerThread(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")
	threadsTable := os.Getenv("DB_THREADS_TABLE")
	usersTable := os.Getenv("DB_USERS_TABLE")     // Assuming you have this environment variable
	repliesTable := os.Getenv("DB_REPLIES_TABLE") // Assuming you have this environment variable

	// Retrieve threadID from URL parameter
	threadID := chi.URLParam(r, "threadID")

	// SQL query to get the thread with join to fetch username and reply count
	threadQuery := fmt.Sprintf(`
        SELECT t.id, t.subforum_id, t.title, t.content, t.created_by, u.name AS created_by_name, t.is_pinned, t.created_at, t.updated_at, COUNT(r.id) AS reply_count
        FROM %s t
        INNER JOIN %s u ON t.created_by = u.id
        LEFT JOIN %s r ON t.id = r.thread_id
        WHERE t.id = $1
        GROUP BY t.id, u.name
    `, threadsTable, usersTable, repliesTable)

	row := database.GetDB().QueryRow(threadQuery, threadID)

	var thread models.Thread
	err := row.Scan(&thread.ID, &thread.SubforumID, &thread.Title, &thread.Content, &thread.CreatedBy, &thread.CreatedByName, &thread.IsPinned, &thread.CreatedAt, &thread.UpdatedAt, &thread.ReplyCount)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving thread: \n%v", err))
		return
	}

	// Respond with the thread using the RespondWithJSON function
	util.RespondWithJSON(w, http.StatusOK, thread)
}
