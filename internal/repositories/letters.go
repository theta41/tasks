package repositories

import "main/internal/models"

type Letters interface {
	GetLettersByTaskName(taskName string) ([]models.Letter, error)
}
