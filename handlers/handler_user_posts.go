package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
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

	// Retrieve the claims from the middleware context (util.AuthenticateUserMiddleware)
	claims, ok := r.Context().Value("userClaims").(*jwt.RegisteredClaims)
	if !ok {
		util.RespondWithError(w, http.StatusInternalServerError, "Error processing user data")
		return
	}

	// SQL query to get the user's threads
	query := fmt.Sprintf(`
	SELECT t.id, t.subforum_id, t.title, t.content, t.created_by, u.name, t.is_pinned, t.created_at, t.updated_at, COUNT(r.id) AS reply_count
	FROM %s t
	INNER JOIN %s u ON t.created_by = u.id
	LEFT JOIN %s r ON t.id = r.thread_id
	WHERE t.created_by = $1
	GROUP BY t.id, u.name
	ORDER BY t.id DESC
	`, threadsTable, usersTable, repliesTable)

	// Execute sql query and return a rows result set
	rows, err := database.GetDB().Query(query, claims.Subject)
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

	// Modify the query to join with users table to get createdByName
	query = fmt.Sprintf(`
        SELECT r.id, r.thread_id, r.content, r.created_by, u.name AS created_by_name, r.created_at, r.updated_at
        FROM %s r
        INNER JOIN %s u ON r.created_by = u.id
        WHERE r.created_by = $1
				ORDER BY r.id DESC
    `, repliesTable, usersTable)

	// Execute SQL query and return a rows result set
	rows, err = database.GetDB().Query(query, claims.Subject)
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

	log.Println(threads)
	log.Println(replies)

	util.RespondWithJSON(w, http.StatusOK, struct {
		Threads []models.Thread
		Replies []models.Reply
	}{Threads: threads, Replies: replies})
}
