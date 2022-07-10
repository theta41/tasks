package pg

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"gitlab.com/g6834/team41/tasks/internal/models"
	"testing"
	"time"
)

func TestTask_AddAndGetTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	tasks := NewTasks(db)

	task := models.Task{
		ID:           1,
		Name:         "test",
		Description:  "test",
		CreatorEmail: "test@example.org",
		CreatedAt:    time.Unix(1656621468, 0),
		FinishedAt:   time.Unix(0, 0),
	}

	mock.ExpectQuery("INSERT INTO tasks (.+) RETURNING id").WithArgs(task.Name, task.Description, task.CreatorEmail, task.CreatedAt.Unix(), task.FinishedAt.Unix()).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	id, err := tasks.AddTask(task)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if id != 1 {
		t.Fatalf("Expected id 1, got %d", id)
	}

	rows := sqlmock.NewRows([]string{"id", "name", "description", "creator_email", "created_at", "finished_at"}).
		AddRow(1, "test", "test", "test@example.org", 1656621468, 0)

	mock.ExpectQuery("SELECT (.+) FROM tasks (.+)").WillReturnRows(rows)
	got, err := tasks.GetTaskByName(task.Name)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if diff := deep.Equal(got, task); diff != nil {
		t.Fatalf("Expected %v\n got %v\n diff: %v\n", task, got, diff)
	}
}
