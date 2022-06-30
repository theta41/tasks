package pg

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/google/uuid"
	"gitlab.com/g6834/team41/tasks/internal/models"
	"testing"
	"time"
)

func TestLetters_GetLettersByTaskName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "email", "order", "task_id", "sent", "answered", "accepted", "accept_uuid", "accepted_at", "sent_at"}).
		AddRow(1, "test@example.org", 1, 1, true, false, false, "00000000-0000-0000-0000-000000000000", 0, 1656621468)

	mock.ExpectQuery("SELECT (.+) FROM letters LEFT JOIN tasks t ON t.id = letters.task_id (.+)").WillReturnRows(rows)

	exp := []models.Letter{{1, "test@example.org", 1, 1, true, false, false, uuid.MustParse("00000000-0000-0000-0000-000000000000"), time.Unix(0, 0), time.Unix(1656621468, 0)}}
	letters := NewLetters(db)
	got, err := letters.GetLettersByTaskName("test")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if diff := deep.Equal(got, exp); diff != nil {
		t.Fatalf("Expected %v\n got %v\n diff: %v\n", exp, got, diff)
	}
}
