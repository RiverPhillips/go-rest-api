package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo"
	"github.com/RiverPhillips/go-rest-api/pkg/resources"
	"github.com/go-chi/chi/v5"
)

var _ resources.Handler = (*todoHandlerV1)(nil)

type todoHandlerV1 struct {
}

func NewTodoHandlerV1() resources.Handler {
	return &todoHandlerV1{}
}

func (t *todoHandlerV1) Route() func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", t.createTodo)
	}
}

func (t *todoHandlerV1) createTodo(w http.ResponseWriter, r *http.Request) {
	var req todo.CreateTodoRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	created := todo.CreateTodoResponse{
		Data: todo.CreateTodoResponseData{
			Type: "todos",
			Attributes: todo.CreateTodoResponseAttributes{
				Title:       req.Data.Attributes.Title,
				Description: req.Data.Attributes.Description,
				Completed:   false,
				UpdatedAt:   time.Now(),
			},
			Id: 1,
			Links: resources.Links{
				Self: fmt.Sprintf("http://localhost:8080/%s/%d", r.URL.Path, 1),
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", created.Data.Links.Self)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
