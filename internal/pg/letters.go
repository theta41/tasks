package pg

import (
	"database/sql"
	"gitlab.com/g6834/team41/tasks/internal/models"
	"time"
)

type Letters struct {
	db *sql.DB
}

func NewLetters(db *sql.DB) *Letters {
	return &Letters{db: db}
}

func (l Letters) GetLettersByTaskName(taskName string) ([]models.Letter, error) {
	letters := make([]models.Letter, 0)
	rows, err := l.db.Query("SELECT id, email, 'order', task_id, sent, answered, accepted, accept_uuid, accepted_at, sent_at FROM letters LEFT JOIN tasks t ON t.id = letters.task_id WHERE t.name = $1", taskName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var letter models.Letter
		var acceptedAt, sentAt int
		if err := rows.Scan(&letter.ID, &letter.Email, &letter.Order, &letter.TaskId, &letter.Sent, &letter.Answered, &letter.Accepted, &letter.AcceptUuid, &acceptedAt, &sentAt); err != nil {
			return nil, err
		}
		letter.AcceptedAt = time.Unix(int64(acceptedAt), 0)
		letter.SentAt = time.Unix(int64(sentAt), 0)
		letters = append(letters, letter)
	}

	return letters, nil
}
