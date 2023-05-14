package repository

import (
	"context"

	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo"
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Todo interface {
	Create(ctx context.Context, todo *todo.CreateTodoAttributes) (todoId int, createdAt pgtype.Timestamp, err error)
}

const (
	InsertIntoTodo = "INSERT INTO todo (title, description, completed) VALUES ($1, $2, $3) RETURNING id, created_at;"
)

var _ Todo = (*todoRepository)(nil)

type todoRepository struct {
	pool *pgxpool.Pool
}

func NewTodoRepository(pool *pgxpool.Pool) Todo {
	return &todoRepository{pool}
}

func (r *todoRepository) Create(ctx context.Context, todo *todo.CreateTodoAttributes) (todoId int, createdAt pgtype.Timestamp, err error) {
	if err = r.pool.QueryRow(ctx, InsertIntoTodo, todo.Title, todo.Description, todo.Completed).Scan(&todoId, &createdAt); err != nil {
		return 0, pgtype.Timestamp{}, err
	}

	return todoId, createdAt, nil
}
