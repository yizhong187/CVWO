package handlers

import (
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

// HandlerAllReplies handles the request to retrieve all replies of a thread from the database.
func HandlerAllReplies(w http.ResponseWriter, r *http.Request) {

	godotenv.Load(".env")
	usersTable := os.Getenv("DB_USERS_TABLE")
	if usersTable == "" {
		log.Fatal("usersTable is not set in the environment")
	}
	threadsTable := os.Getenv("DB_THREADS_TABLE")
	if usersTable == "" {
		log.Fatal("threadsTable is not set in the environment")
	}
	repliesTable := os.Getenv("DB_REPLIES_TABLE")
	if usersTable == "" {
		log.Fatal("repliesTable is not set in the environment")
	}
	subforumsTable := os.Getenv("DB_SUBFORUMS_TABLE")
	if subforumsTable == "" {
		log.Fatal("subforumTable is not set in the environment")
	}

	threadID := chi.URLParam(r, "threadID")

	query := fmt.Sprintf(`
	SELECT r.id, r.thread_id, t.title AS thread_name, t.subforum_id, sf.name AS subforum_name, r.content, r.created_by, u.name AS created_by_name, r.created_at, r.updated_at
	FROM %s r
	INNER JOIN %s u ON r.created_by = u.id
	INNER JOIN %s t ON r.thread_id = t.id
	INNER JOIN %s sf ON t.subforum_id = sf.id
	WHERE r.thread_id = $1
	ORDER BY r.id DESC
`, repliesTable, usersTable, threadsTable, subforumsTable)

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
		err := rows.Scan(&reply.ID, &reply.ThreadID, &reply.ThreadName, &reply.SubforumID, &reply.SubforumName,
			&reply.Content, &reply.CreatedBy, &reply.CreatedByName, &reply.CreatedAt, &reply.UpdatedAt)
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
