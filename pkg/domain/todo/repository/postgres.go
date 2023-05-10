package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo"
)

type Todo interface {
	Create(ctx context.Context, todo *todo.CreateTodoAttributes) (int, time.Time, error)
}

const (
	InsertIntoTodo = "INSERT INTO todo (title, description, completed) VALUES ($1, $2, $3) RETURNING id, created_at;"
)

type todoRepository struct {
	db sql.DB
}

func NewTodoRepository(db sql.DB) Todo {
	return &todoRepository{db}
}

func (r *todoRepository) Create(ctx context.Context, todo *todo.CreateTodoAttributes) (todoId int, createdAt time.Time, err error) {
	if err = r.db.QueryRowContext(ctx, InsertIntoTodo, todo.Title, todo.Description, todo.Completed).Scan(&todoId, createdAt); err != nil {
		return 0, time.Time{}, err
	}

	return todoId, createdAt, nil
}
