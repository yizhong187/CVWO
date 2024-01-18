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

// HandlerAllReplies handles the request to retrieve all replies of a thread from the database.
func HandlerAllReplies(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	repliesTable := os.Getenv("DB_REPLIES_TABLE")
	usersTable := os.Getenv("DB_USERS_TABLE") // Assuming you have this environment variable

	// Modify the query to join with users table to get createdByName
	query := fmt.Sprintf(`
        SELECT r.id, r.thread_id, r.content, r.created_by, u.name AS created_by_name, r.created_at, r.updated_at
        FROM %s r
        INNER JOIN %s u ON r.created_by = u.id
        WHERE r.thread_id = $1
				ORDER BY r.id DESC
    `, repliesTable, usersTable)

	threadID := chi.URLParam(r, "threadID")

	// Execute SQL query and return a rows result set
	rows, err := database.GetDB().Query(query, threadID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}
	defer rows.Close()

	// Initialize a slice of replies, scan each row into a reply struct, and append it to the slice
	var replies []models.Reply
	for rows.Next() {
		var reply models.Reply
		err := rows.Scan(&reply.ID, &reply.ThreadID, &reply.Content, &reply.CreatedBy, &reply.CreatedByName, &reply.CreatedAt, &reply.UpdatedAt)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error scanning row: \n%v", err))
			return
		}
		replies = append(replies, reply)
	}

	if err = rows.Err(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Row error: \n%v", err))
		return
	}

	// Respond with the list of replies in JSON format
	util.RespondWithJSON(w, http.StatusOK, replies)
}
