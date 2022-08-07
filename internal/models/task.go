package models

import "time"

type Task struct {
	ID           int       `json:"id" example:"123"`
	Name         string    `json:"name" example:"Test task"`
	Description  string    `json:"description" example:"Description task"`
	CreatorEmail string    `json:"creator_email" example:"creator@task.com"`
	CreatedAt    time.Time `json:"created_at" example:"2021-05-25T00:53:16.535668Z" format:"date-time"`
	FinishedAt   time.Time `json:"ended_at" example:"2021-05-25T00:53:16.535668Z" format:"date-time"`
}
