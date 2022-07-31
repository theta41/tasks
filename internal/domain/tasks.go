package domain

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"gitlab.com/g6834/team41/tasks/internal/models"
)

func CreateTask(task models.Task, emails []string) error {
	id, err := env.E.TR.AddTask(task)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	// Send task to kafka
	go func() {
		err = env.E.K.Publish([]byte("create-task"), []byte(fmt.Sprintf(`{"task_id": %d}`, id)))
		if err != nil {
			logrus.Error("env.E.K.Publish error: ", err)
			sentry.CaptureException(err)
		}
	}()

	for i := range emails {
		l := models.Letter{
			Email:      emails[i],
			Order:      i,
			TaskId:     id,
			Sent:       false,
			Answered:   false,
			Accepted:   false,
			AcceptUuid: uuid.New(),
			AcceptedAt: time.Unix(0, 0),
			SentAt:     time.Now(),
		}
		// TODO: batch letters.
		err = env.E.LR.AddLetter(l)
		if err != nil {
			return fmt.Errorf("failed to add letter: %w", err)
		}

		// Send task to kafka
		go func() {
			err = env.E.K.Publish([]byte("create-letter"), []byte(fmt.Sprintf(`{"email": "%s", "task_id": %d}`, l.Email, l.TaskId)))
			if err != nil {
				logrus.Error("env.E.K.Publish error: ", err)
				sentry.CaptureException(err)
			}
		}()
	}

	// TODO: Send emails
	// TODO: Update sent status
	return nil
}

func DeleteTask(taskId int) error {
	err := env.E.TR.DeleteTaskById(taskId)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	// Send delete to kafka
	go func() {
		err = env.E.K.Publish([]byte("delete"), []byte(strconv.Itoa(taskId)))
		if err != nil {
			logrus.Error("env.E.K.Publish error: ", err)
			sentry.CaptureException(err)
		}
	}()

	return nil
}

func ListTasks(email string) ([]models.Task, error) {
	tasks, err := env.E.TR.GetAllTasksByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks: %w", err)
	}
	return tasks, nil
}

func GetTask(id int) (*models.Task, []models.Letter, error) {
	task, err := env.E.TR.GetTask(id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get task: %w", err)
	}

	letters, err := env.E.LR.GetLettersByTaskId(id)
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
