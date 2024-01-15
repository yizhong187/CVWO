package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/yizhong187/CVWO/database"
	"github.com/yizhong187/CVWO/models"
	"github.com/yizhong187/CVWO/util"
)

func HandlerThread(w http.ResponseWriter, r *http.Request) {
	godotenv.Load(".env")
	threadsTable := os.Getenv("DB_THREADS_TABLE")
	repliesTable := os.Getenv("DB_REPLIES_TABLE")

	// Retrieve threadID from URL parameter
	threadID := chi.URLParam(r, "threadID")

	// First SQL query to get the thread
	threadQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", threadsTable)
	row := database.GetDB().QueryRow(threadQuery, threadID)
	var thread models.Thread
	err := row.Scan(&thread.ID, &thread.SubforumID, &thread.Title, &thread.Content, &thread.CreatedBy, &thread.IsPinned, &thread.CreatedAt, &thread.UpdatedAt)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving thread: \n%v", err))
		return
	}

	// Second SQL query to get replies associated with the thread
	repliesQuery := fmt.Sprintf("SELECT * FROM %s WHERE thread_id = $1", repliesTable)
	rows, err := database.GetDB().Query(repliesQuery, threadID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving replies: \n%v", err))
		return
	}
	defer rows.Close()

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

	// Third SQL query to get replyCount for the thread
	numThreadID, err := strconv.Atoi(threadID)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid Parameter: \n%v", err))
		return
	}

	replyCount, err := util.QueryReplyCount(numThreadID)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error retrieving replycount: \n%v", err))
		return
	}

	// Respond with thread and its replies using the RespondWithJSON function
	response := struct {
		Thread     models.Thread  `json:"thread"`
		Replies    []models.Reply `json:"replies"`
		ReplyCount int            `json:"replyCount"`
	}{
		Thread:     thread,
		Replies:    replies,
		ReplyCount: replyCount,
	}
	util.RespondWithJSON(w, http.StatusOK, response)
}
