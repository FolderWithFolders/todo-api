package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

// Подключение к базе данных
func ConnectDB() error {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	db = conn
	return nil
}

// Создать задачу
func CreateTask(task *Task) error {
	query := `
        INSERT INTO tasks (title, description, status)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `
	return db.QueryRow(
		context.Background(),
		query,
		task.Title,
		task.Description,
		task.Status,
	).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

// Получить все задачи
func GetTasks() ([]Task, error) {
	rows, err := db.Query(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// Обновить задачу
func UpdateTask(id int, task *Task) error {
	query := `
        UPDATE tasks
        SET title = $1, description = $2, status = $3, updated_at = now()
        WHERE id = $4
        RETURNING updated_at
    `
	return db.QueryRow(
		context.Background(),
		query,
		task.Title,
		task.Description,
		task.Status,
		id,
	).Scan(&task.UpdatedAt)
}

// Удалить задачу
func DeleteTask(id int) error {
	_, err := db.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	return err
}
