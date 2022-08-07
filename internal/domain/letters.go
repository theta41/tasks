package domain

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"gitlab.com/g6834/team41/tasks/internal/kafka"
	"gitlab.com/g6834/team41/tasks/internal/models"
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
		err = env.E.K.PublishAnalytics([]byte("accept"), []byte(fmt.Sprintf(`{"task_id": %d, "email": "%s"}`, task.ID, letter.Email)))
		if err != nil {
			logrus.Error("env.E.K.PublishAnalytics error: ", err)
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

			// Send email to next letter
			sendAcceptanceEmail(letters[i])

			found = true
			break
		}
	}
	if !found {
		go func() {
			err = env.E.K.PublishAnalytics([]byte("finish"), []byte(fmt.Sprintf(`{"task_id": %d}`, task.ID)))
			if err != nil {
				logrus.Error("env.E.K.PublishAnalytics error: ", err)
				sentry.CaptureException(err)
			}
		}()
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
		err = env.E.K.PublishAnalytics([]byte("decline"), []byte(fmt.Sprintf(`{"task_id": %d, "email": "%s"}`, letter.TaskId, letter.Email)))
		if err != nil {
			logrus.Error("env.E.K.PublishAnalytics error: ", err)
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

		if letters[i].Order < letter.Order {
			// Send email
			sendСancellationEmail(letters[i])
		}
	}

	return nil
}

func sendAcceptanceEmail(l models.Letter) {
	baseUrl := fmt.Sprintf("http://tasks%v", env.E.C.HostAddress)

	key := []byte("email")
	value, err := kafka.MakeAcceptanceEmail(l.Email, baseUrl, l.TaskId, l.AcceptUuid)
	if err != nil {
		logrus.Error("MakeAcceptanceEmail error: ", err)
		sentry.CaptureException(err)
		return
	}

	go func() {
		err := env.E.K.PublishEmail(key, value)
		if err != nil {
			logrus.Error("env.E.K.PublishEmail error: ", err)
			sentry.CaptureException(err)
		}
	}()
}

func sendСancellationEmail(l models.Letter) {
	key := []byte("email")
	value, err := kafka.MakeCancellationEmail(l.Email, l.TaskId)
	if err != nil {
		logrus.Error("MakeCancellationEmail error: ", err)
		sentry.CaptureException(err)
		return
	}

	go func() {
		err := env.E.K.PublishEmail(key, value)
		if err != nil {
			logrus.Error("env.E.K.PublishEmail error: ", err)
			sentry.CaptureException(err)
		}
	}()
}
