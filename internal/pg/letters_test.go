package pg

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/google/uuid"
	"gitlab.com/g6834/team41/tasks/internal/models"
)

func TestLetters_GetLettersByTaskId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "order", "task_id", "sent", "answered", "accepted", "accept_uuid", "accepted_at", "sent_at"}).
		AddRow(1, "test@example.org", 1, 1, true, false, false, "00000000-0000-0000-0000-000000000000", 0, 1656621468)

	mock.ExpectQuery("SELECT (.+) FROM letters (.+)").WillReturnRows(rows)

	exp := []models.Letter{{
		ID:         1,
		Email:      "test@example.org",
		Order:      1,
		TaskId:     1,
		Sent:       true,
		Answered:   false,
		Accepted:   false,
		AcceptUuid: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		AcceptedAt: time.Unix(0, 0),
		SentAt:     time.Unix(1656621468, 0),
	}}
	letters := NewLetters(db)
	got, err := letters.GetLettersByTaskId(1)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if diff := deep.Equal(got, exp); diff != nil {
		t.Fatalf("Expected %v\n got %v\n diff: %v\n", exp, got, diff)
	}
}
