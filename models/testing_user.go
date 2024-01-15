package models

import "time"

type TestingUser struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"-"`
	Type         string    `json:"type"`
	CreatedAt    time.Time `json:"createdAt"`
}
