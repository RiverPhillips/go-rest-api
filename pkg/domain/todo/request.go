package todo

import (
	"github.com/RiverPhillips/go-rest-api/pkg/resources"
)

// CreateTodoRequest is the request for creating a todo
type CreateTodoRequest = resources.CreateRequest[CreateTodoAttributes]

// CreateTodoAttributes are the attributes for creating a todo
type CreateTodoAttributes struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
