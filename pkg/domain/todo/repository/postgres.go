// Package repository is the package that contains the repository interfaces and implementations
package repository

import (
	"context"

	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo"
	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Todo is the interface for the todo repository
type Todo interface {
	Create(ctx context.Context, todo *todo.CreateTodoAttributes) (todoID int, createdAt pgtype.Timestamp, err error)
}

const (
	insertIntoTodo = "INSERT INTO todo (title, description, completed) VALUES ($1, $2, $3) RETURNING id, created_at;"
)

var _ Todo = (*todoRepository)(nil)

type todoRepository struct {
	pool *pgxpool.Pool
}

// NewTodoRepository creates a new instance of TodoRepository
func NewTodoRepository(pool *pgxpool.Pool) Todo {
	return &todoRepository{pool}
}

// Create creates a new todo in the database
func (r *todoRepository) Create(ctx context.Context, todo *todo.CreateTodoAttributes) (todoID int, createdAt pgtype.Timestamp, err error) {
	if err = r.pool.QueryRow(ctx, insertIntoTodo, todo.Title, todo.Description, todo.Completed).Scan(&todoID, &createdAt); err != nil {
		return 0, pgtype.Timestamp{}, err
	}

	return todoID, createdAt, nil
}
