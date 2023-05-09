package todo

import (
	"github.com/RiverPhillips/go-rest-api/pkg/resources"
)

type CreateTodoRequest = resources.CreateRequest[CreateTodoAttributes]

type CreateTodoAttributes struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
