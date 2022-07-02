package domain

import (
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"gitlab.com/g6834/team41/tasks/internal/models"
	"time"
)

func CreateTask(task models.Task, emails []string) error {
	id, err := env.E.TR.AddTask(task)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	for i := range emails {
		l := models.Letter{
			Email:      emails[i],
			Order:      i,
			TaskId:     id,
			Sent:       false,
			Answered:   false,
			Accepted:   false,
			AcceptUuid: uuid.New(),
			AcceptedAt: time.Time{},
			SentAt:     time.Now(),
		}
		// TODO: batch letters.
		err = env.E.LR.AddLetter(l)
		if err != nil {
			return fmt.Errorf("failed to add letter: %w", err)
		}
	}

	// TODO: Send emails
	// TODO: Update sent status
	return nil
}

func DeleteTask(name string) error {
	err := env.E.TR.DeleteTaskByName(name)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

func ListTasks(email string) ([]models.Task, error) {
	tasks, err := env.E.TR.GetAllTasksByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	return tasks, nil
}

func GetTask(name string) (*models.Task, []models.Letter, error) {
	task, err := env.E.TR.GetTaskByName(name)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get task: %w", err)
	}

	letters, err := env.E.LR.GetLettersByTaskName(name)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get letters: %w", err)
	}

	return &task, letters, nil
}

func UpdateTask(task models.Task) error {
	err := env.E.TR.UpdateTask(task)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}
