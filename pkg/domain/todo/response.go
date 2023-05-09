package todo

import (
	"time"

	"github.com/RiverPhillips/go-rest-api/pkg/resources"
)

type CreateTodoResponse = resources.CreatedResponse[CreateTodoResponseAttributes, int]
type CreateTodoResponseData = resources.CreateResponseData[CreateTodoResponseAttributes, int]

type CreateTodoResponseAttributes struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
