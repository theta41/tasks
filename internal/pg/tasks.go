package pg

import (
	"database/sql"
	"gitlab.com/g6834/team41/tasks/internal/models"
	"time"
)

type Task struct {
	db *sql.DB
}

func NewTask(db *sql.DB) Task {
	return Task{db: db}
}

func (t Task) AddTask(task models.Task) (int, error) {
	row := t.db.QueryRow("INSERT INTO tasks (name, description, creator_email, created_at, finished_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		task.Name, task.Description, task.CreatorEmail, task.CreatedAt, task.FinishedAt)
	if row.Err() != nil {
		return 0, row.Err()
	}

	err := row.Scan(&task.ID)
	if err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (t Task) DeleteTaskByName(name string) error {
	_, err := t.db.Exec("DELETE FROM tasks WHERE name = $1", name)
	return err
}

func (t Task) GetTaskByName(name string) (models.Task, error) {
	var task models.Task
	var createdAt, finishedAt int64
	err := t.db.QueryRow("SELECT id, name, description, creator_email, created_at, finished_at FROM tasks WHERE name = $1", name).
		Scan(&task.ID, &task.Name, &task.Description, &task.CreatorEmail, &createdAt, &finishedAt)
	task.CreatedAt = time.Unix(createdAt, 0)
	task.FinishedAt = time.Unix(finishedAt, 0)
	return task, err
}

func (t Task) GetAllTasksByEmail(email string) ([]models.Task, error) {
	tasks := make([]models.Task, 0)
	rows, err := t.db.Query("SELECT id, name, description, creator_email, created_at, finished_at FROM tasks WHERE creator_email = $1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.CreatorEmail, &task.CreatedAt, &task.FinishedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t Task) UpdateTask(task models.Task) error {
	_, err := t.db.Exec("UPDATE tasks SET name = $1, description = $2, creator_email = $3, created_at = $4, finished_at = $5 WHERE id = $6",
		task.Name, task.Description, task.CreatorEmail, task.CreatedAt, task.FinishedAt, task.ID)
	return err
}
