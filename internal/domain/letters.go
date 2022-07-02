package domain

import (
	"fmt"
	"github.com/google/uuid"
	"gitlab.com/g6834/team41/tasks/internal/env"
	"time"
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

	// Find next letter to send email or complete agreement.
	letters, err := env.E.LR.GetLettersByTaskName(task.Name)
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

	task, err := env.E.TR.GetTask(letter.TaskId)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Send cancel notification to all previous letters and creator.
	letters, err := env.E.LR.GetLettersByTaskName(task.Name)
	for i := range letters {
		if letters[i].Order < letter.Order {
			// TODO: Send email.
		}
	}

	return nil
}
