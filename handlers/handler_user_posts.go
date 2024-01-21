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

func HandlerUserPosts(w http.ResponseWriter, r *http.Request) {

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

	// Retrieve username from URL query
	userName := chi.URLParam(r, "userName")
	if userName == "" {
		util.RespondWithError(w, http.StatusBadRequest, "userName is required")
		return
	}

	// SQL query to get the user's threads
	threadQuery := fmt.Sprintf(`
	SELECT t.id, t.subforum_id, sf.name AS subforum_name, t.title, t.content, t.created_by, u.name, t.is_pinned, t.created_at, t.updated_at, COUNT(r.id) AS reply_count
	FROM %s t
	INNER JOIN %s u ON t.created_by = u.id
	LEFT JOIN %s r ON t.id = r.thread_id
	INNER JOIN %s sf ON t.subforum_id = sf.id
	WHERE u.name = $1
	GROUP BY t.id, u.name, sf.name
	ORDER BY t.id DESC
`, threadsTable, usersTable, repliesTable, subforumsTable)

	// Execute sql query and return a rows result set for threads
	threadRows, err := database.GetDB().Query(threadQuery, userName)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}
	defer threadRows.Close()

	// Initialize a slice of threads, scan each row into a thread struct and append into the slice
	var threads []models.Thread
	for threadRows.Next() {
		var thread models.Thread
		err := threadRows.Scan(&thread.ID, &thread.SubforumID, &thread.SubforumName, &thread.Title, &thread.Content, &thread.CreatedBy,
			&thread.CreatedByName, &thread.IsPinned, &thread.CreatedAt, &thread.UpdatedAt, &thread.ReplyCount)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error scanning row: \n%v", err))
			return
		}
		threads = append(threads, thread)
	}

	if err = threadRows.Err(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Row error: \n%v", err))
		return
	}

	// SQL query to get thread and subforum info for replies
	replyQuery := fmt.Sprintf(`
	SELECT r.id, r.thread_id, t.title AS thread_name, t.subforum_id, sf.name AS subforum_name, r.content, r.created_by, u.name AS created_by_name, r.created_at, r.updated_at
	FROM %s r
	INNER JOIN %s u ON r.created_by = u.id
	INNER JOIN %s t ON r.thread_id = t.id
	INNER JOIN %s sf ON t.subforum_id = sf.id
	WHERE u.name = $1
	ORDER BY r.id DESC
`, repliesTable, usersTable, threadsTable, subforumsTable)

	// Execute SQL query and return a rows result set for replies
	replyRows, err := database.GetDB().Query(replyQuery, userName)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Internal Server Error: \n%v", err))
		return
	}
	defer replyRows.Close()

	// Initialize a slice of replies, scan each row into a reply struct, and append it to the slice
	var replies []models.Reply
	for replyRows.Next() {
		var reply models.Reply
		err := replyRows.Scan(&reply.ID, &reply.ThreadID, &reply.ThreadName, &reply.SubforumID, &reply.SubforumName, &reply.Content, &reply.CreatedBy, &reply.CreatedByName, &reply.CreatedAt, &reply.UpdatedAt)
		if err != nil {
			util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error scanning row: \n%v", err))
			return
		}
		replies = append(replies, reply)
	}

	if err = replyRows.Err(); err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Row error: \n%v", err))
		return
	}

	util.RespondWithJSON(w, http.StatusOK, struct {
		Threads []models.Thread `json:"threads"`
		Replies []models.Reply  `json:"replies"`
	}{Threads: threads, Replies: replies})
}
