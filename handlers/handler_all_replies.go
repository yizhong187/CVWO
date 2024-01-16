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
	query := fmt.Sprintf("SELECT * FROM %s WHERE thread_id = $1", repliesTable)

	threadID := chi.URLParam(r, "threadID")

	// Execute sql query and return a rows result set
	rows, err := database.GetDB().Query(query, threadID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}
	defer rows.Close()

	// Initialize a slice of threads, scan each row into a thread struct and append into the slice
	var replies []models.Reply
	for rows.Next() {
		var reply models.Reply
		err := rows.Scan(&reply.ID, &reply.ThreadID, &reply.Content, &reply.CreatedBy, &reply.CreatedAt, &reply.UpdatedAt)
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

	// Respond with the list of threads in JSON format
	util.RespondWithJSON(w, http.StatusOK, replies)
}
