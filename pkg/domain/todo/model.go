// Package todo is the package that contains the todo domain
package todo

// Schema is the database schema for the todo table
type Schema struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Completed   bool   `db:"completed"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
