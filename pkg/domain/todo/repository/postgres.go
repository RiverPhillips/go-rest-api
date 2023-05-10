package repository

import (
	"context"
	"time"

	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Todo interface {
	Create(ctx context.Context, todo *todo.CreateTodoAttributes) (int, time.Time, error)
}

const (
	InsertIntoTodo = "INSERT INTO todo (title, description, completed) VALUES ($1, $2, $3) RETURNING id, created_at;"
)

type todoRepository struct {
	pool *pgxpool.Pool
}

func NewTodoRepository(pool *pgxpool.Pool) Todo {
	return &todoRepository{pool}
}

func (r *todoRepository) Create(ctx context.Context, todo *todo.CreateTodoAttributes) (todoId int, createdAt time.Time, err error) {
	if err = r.pool.QueryRow(ctx, InsertIntoTodo, todo.Title, todo.Description, todo.Completed).Scan(&todoId, createdAt); err != nil {
		return 0, time.Time{}, err
	}

	return todoId, createdAt, nil
}
