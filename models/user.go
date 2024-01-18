package models

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"-"`
	Type         string    `json:"type"`
	CreatedAt    time.Time `json:"createdAt"`
}
