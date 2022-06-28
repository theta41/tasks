package repositories

import "gitlab.com/g6834/team41/tasks/internal/models"

type Letters interface {
	GetLettersByTaskName(taskName string) ([]models.Letter, error)
}
