package repositories

import "gitlab.com/g6834/team41/tasks/internal/models"

type Letters interface {
	AddLetter(letter models.Letter) error
	UpdateLetter(letter models.Letter) error
	GetLettersByTaskName(taskName string) ([]models.Letter, error)
}
