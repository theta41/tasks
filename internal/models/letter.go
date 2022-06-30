package models

import (
	"github.com/google/uuid"
	"time"
)

type Letter struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Order      int       `json:"order"`
	TaskId     int       `json:"task_id"`
	Sent       bool      `json:"sent"`
	Answered   bool      `json:"answered"`
	Accepted   bool      `json:"accepted"`
	AcceptUuid uuid.UUID `json:"accept_uuid"`
	AcceptedAt time.Time `json:"accept_at"`
	SentAt     time.Time `json:"sent_at"`
}
