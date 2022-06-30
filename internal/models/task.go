package models

import "time"

type Task struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreatorEmail string    `json:"creator_email"`
	CreatedAt    time.Time `json:"created_at"`
	FinishedAt   time.Time `json:"ended_at"`
}
