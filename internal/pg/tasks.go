package pg

import (
	"database/sql"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/g6834/team41/tasks/internal/models"
)

type Tasks struct {
	db *sql.DB
}

func NewTasks(db *sql.DB) Tasks {
	return Tasks{db: db}
}

func (t Tasks) AddTask(task models.Task) (int, error) {
	row := t.db.QueryRow("INSERT INTO tasks (name, description, creator_email, created_at, finished_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		task.Name, task.Description, task.CreatorEmail, task.CreatedAt.Unix(), task.FinishedAt.Unix())
	if row.Err() != nil {
		return 0, row.Err()
	}

	err := row.Scan(&task.ID)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (t Tasks) DeleteTaskByName(name string) error {
	_, err := t.db.Exec("DELETE FROM tasks WHERE name = $1", name)
	return err
}

func (t Tasks) GetTaskByName(name string) (models.Task, error) {
	var task models.Task
	var createdAt, finishedAt int64
	err := t.db.QueryRow("SELECT id, name, description, creator_email, created_at, finished_at FROM tasks WHERE name = $1", name).
		Scan(&task.ID, &task.Name, &task.Description, &task.CreatorEmail, &createdAt, &finishedAt)
	task.CreatedAt = time.Unix(createdAt, 0)
	task.FinishedAt = time.Unix(finishedAt, 0)
	return task, err
}

func (t Tasks) GetAllTasksByEmail(email string) ([]models.Task, error) {
	tasks := make([]models.Task, 0)
	rows, err := t.db.Query("SELECT id, name, description, creator_email, created_at, finished_at FROM tasks WHERE creator_email = $1", email)
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
		var task models.Task
		var createdAt, finishedAt int64
		if err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.CreatorEmail, &createdAt, &finishedAt); err != nil {
			return nil, err
		}
		task.CreatedAt = time.Unix(createdAt, 0)
		task.FinishedAt = time.Unix(finishedAt, 0)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t Tasks) UpdateTask(task models.Task) error {
	_, err := t.db.Exec("UPDATE tasks SET name = $1, description = $2, creator_email = $3, created_at = $4, finished_at = $5 WHERE id = $6",
		task.Name, task.Description, task.CreatorEmail, task.CreatedAt.Unix(), task.FinishedAt.Unix(), task.ID)
	return err
}

func (t Tasks) GetTask(id int) (models.Task, error) {
	var task models.Task
	var createdAt, finishedAt int64
	err := t.db.QueryRow("SELECT id, name, description, creator_email, created_at, finished_at FROM tasks WHERE id = $1", id).
		Scan(&task.ID, &task.Name, &task.Description, &task.CreatorEmail, &createdAt, &finishedAt)
	task.CreatedAt = time.Unix(createdAt, 0)
	task.FinishedAt = time.Unix(finishedAt, 0)
	return task, err
}
