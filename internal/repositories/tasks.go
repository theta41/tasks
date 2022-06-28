package repositories

import "main/internal/models"

type Tasks interface {
	AddTask(task models.Task) error
	DeleteTaskByName(name string) error
	GetTaskByName(name string) (models.Task, error)
	GetTasksByEmail(email string) ([]models.Task, error)
	UpdateTask(task models.Task) error
}
