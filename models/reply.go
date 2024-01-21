package models

import "time"

type Reply struct {
	ID            int       `json:"id"`
	ThreadID      int       `json:"threadID"`
	Content       string    `json:"content"`
	CreatedBy     string    `json:"createdBy"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	CreatedByName string    `json:"createdByName"`
	ThreadName    string    `json:"threadName"`
	SubforumID    int       `json:"subforumID"`
	SubforumName  string    `json:"subforumName"`
}
