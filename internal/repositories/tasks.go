package repositories

import "gitlab.com/g6834/team41/tasks/internal/models"

type Tasks interface {
	AddTask(task models.Task) error
	DeleteTaskByName(name string) error
	GetTaskByName(name string) (models.Task, error)
	GetTasksByEmail(email string) ([]models.Task, error)
	UpdateTask(task models.Task) error
}
