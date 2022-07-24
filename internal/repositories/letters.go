package repositories

import "gitlab.com/g6834/team41/tasks/internal/models"

type Letters interface {
	AddLetter(letter models.Letter) error
	UpdateLetter(letter models.Letter) error
	GetLetterByUUID(uuid string) (models.Letter, error)
	GetLettersByTaskId(taskId int) ([]models.Letter, error)
}
