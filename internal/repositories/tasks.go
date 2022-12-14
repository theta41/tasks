package repositories

import "gitlab.com/g6834/team41/tasks/internal/models"

type Tasks interface {
	AddTask(task models.Task) (insertedId int, err error)
	DeleteTaskById(id int) error
	GetTaskByName(name string) (models.Task, error)
	GetAllTasksByEmail(email string) ([]models.Task, error)
	UpdateTask(task models.Task) error
	GetTask(id int) (models.Task, error)
}
