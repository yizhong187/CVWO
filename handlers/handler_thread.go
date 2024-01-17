package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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

	// Retrieve threadID from URL parameter
	threadID := chi.URLParam(r, "threadID")

	// SQL query to get the thread
	threadQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", threadsTable)
	row := database.GetDB().QueryRow(threadQuery, threadID)
	var thread models.Thread
	err := row.Scan(&thread.ID, &thread.SubforumID, &thread.Title, &thread.Content, &thread.CreatedBy, &thread.IsPinned, &thread.CreatedAt, &thread.UpdatedAt)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error retrieving thread: \n%v", err))
		return
	}

	// Used util.QueryReplyCount to get replyCount of the thread
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
		ID         int       `json:"id"`
		SubforumID int       `json:"subforumID"`
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		CreatedBy  string    `json:"createdBy"`
		IsPinned   bool      `json:"isPinned"`
		CreatedAt  time.Time `json:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
		ReplyCount int       `json:"replyCount"`
	}{
		ID:         thread.ID,
		SubforumID: thread.SubforumID,
		Title:      thread.Title,
		Content:    thread.Content,
		CreatedBy:  thread.CreatedBy,
		IsPinned:   thread.IsPinned,
		CreatedAt:  thread.CreatedAt,
		UpdatedAt:  thread.UpdatedAt,
		ReplyCount: replyCount,
	}
	util.RespondWithJSON(w, http.StatusOK, response)
}
