package domain

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/google/uuid"
	"gitlab.com/g6834/team41/tasks/internal/env"
)

func Accept(id uuid.UUID) error {
	letter, err := env.E.LR.GetLetterByUUID(id.String())
	if err != nil {
		return fmt.Errorf("failed to get letter: %w", err)
	}

	letter.Accepted = true
	letter.AcceptedAt = time.Now()

	err = env.E.LR.UpdateLetter(letter)
	if err != nil {
		return fmt.Errorf("failed to update letter: %w", err)
	}

	task, err := env.E.TR.GetTask(letter.TaskId)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Send accept to kafka
	go func() {
		err = env.E.K.Publish([]byte("accept"), []byte(fmt.Sprintf(`{"task_id": %d, "email": "%s"}`, task.ID, letter.Email)))
		if err != nil {
			logrus.Error("env.E.K.Publish error: ", err)
			sentry.CaptureException(err)
		}
	}()

	// Find next letter to send email or complete agreement.
	letters, err := env.E.LR.GetLettersByTaskId(task.ID)
	if err != nil {
		return fmt.Errorf("failed to get letters: %w", err)
	}
	found := false
	for i := range letters {
		if letters[i].Order == letter.Order+1 {
			// TODO: Send email to next letter
			found = true
			break
		}
	}
	if !found {
		go func() {
			err = env.E.K.Publish([]byte("finish"), []byte(fmt.Sprintf(`{"task_id": %d}`, task.ID)))
			if err != nil {
				logrus.Error("env.E.K.Publish error: ", err)
				sentry.CaptureException(err)
			}
		}()
		// TODO: Complete agreement.
	}

	return nil
}

func Decline(id uuid.UUID) error {
	letter, err := env.E.LR.GetLetterByUUID(id.String())
	if err != nil {
		return fmt.Errorf("failed to get letter: %w", err)
	}

	letter.Accepted = false
	letter.AcceptedAt = time.Now()

	err = env.E.LR.UpdateLetter(letter)
	if err != nil {
		return fmt.Errorf("failed to update letter: %w", err)
	}

	// Send decline to kafka
	go func() {
		err = env.E.K.Publish([]byte("decline"), []byte(fmt.Sprintf(`{"task_id": %d, "email": "%s"}`, letter.TaskId, letter.Email)))
		if err != nil {
			logrus.Error("env.E.K.Publish error: ", err)
			sentry.CaptureException(err)
		}
	}()

	task, err := env.E.TR.GetTask(letter.TaskId)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Send cancel notification to all previous letters and creator.
	letters, _ := env.E.LR.GetLettersByTaskId(task.ID)
	for i := range letters {

		//lint:ignore SA9003 coz of todo
		if letters[i].Order < letter.Order {
			// TODO: Send email.
		}
	}

	return nil
}
