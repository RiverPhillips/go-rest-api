package todo

import (
	"time"

	"github.com/RiverPhillips/go-rest-api/pkg/resources"
)

// CreateTodoResponse is the response for creating a todo
type CreateTodoResponse = resources.CreatedResponse[CreateTodoResponseAttributes, int]

// CreateTodoResponseData is the data for creating a todo
type CreateTodoResponseData = resources.CreateResponseData[CreateTodoResponseAttributes, int]

// CreateTodoResponseAttributes are the attributes for creating a todo
type CreateTodoResponseAttributes struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
