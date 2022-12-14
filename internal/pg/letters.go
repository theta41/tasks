package pg

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/models"
)

type Letters struct {
	db *sql.DB
}

func NewLetters(db *sql.DB) *Letters {
	return &Letters{db: db}
}

func (l Letters) GetLettersByTaskId(taskId int) ([]models.Letter, error) {
	letters := make([]models.Letter, 0)
	rows, err := l.db.Query("SELECT id, email, \"order\", task_id, sent, answered, accepted, accept_uuid, accepted_at, sent_at FROM letters WHERE task_id = $1", taskId)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

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

func (l Letters) AddLetter(letter models.Letter) error {
	_, err := l.db.Exec("INSERT INTO letters (email, \"order\", task_id, sent, answered, accepted, accept_uuid, accepted_at, sent_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		letter.Email, letter.Order, letter.TaskId, letter.Sent, letter.Answered, letter.Accepted, letter.AcceptUuid, letter.AcceptedAt.Unix(), letter.SentAt.Unix())
	return err
}

func (l Letters) UpdateLetter(letter models.Letter) error {
	_, err := l.db.Exec("UPDATE letters SET email = $1, \"order\" = $2, task_id = $3, sent = $4, answered = $5, accepted = $6, accept_uuid = $7, accepted_at = $8, sent_at = $9 WHERE id = $10",
		letter.Email, letter.Order, letter.TaskId, letter.Sent, letter.Answered, letter.Accepted, letter.AcceptUuid, letter.AcceptedAt.Unix(), letter.SentAt.Unix(), letter.ID)
	return err
}

func (l Letters) GetLetterByUUID(uuid string) (models.Letter, error) {
	var letter models.Letter
	var acceptedAt, sentAt int
	err := l.db.QueryRow("SELECT id, email, \"order\", task_id, sent, answered, accepted, accept_uuid, accepted_at, sent_at FROM letters WHERE accept_uuid = $1", uuid).
		Scan(&letter.ID, &letter.Email, &letter.Order, &letter.TaskId, &letter.Sent, &letter.Answered, &letter.Accepted, &letter.AcceptUuid, &acceptedAt, &sentAt)
	letter.AcceptedAt = time.Unix(int64(acceptedAt), 0)
	letter.SentAt = time.Unix(int64(sentAt), 0)
	return letter, err
}
