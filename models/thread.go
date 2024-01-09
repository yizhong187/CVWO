package models

import "time"

type Thread struct {
	ID         int       `json:"id"`
	SubforumID int       `json:"subforumID"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedBy  string    `json:"createdBy"`
	IsPinned   bool      `json:"isPinned"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
