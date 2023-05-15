// Package handler is the package that contains the http handlers for the todo domain
package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo"
	"github.com/RiverPhillips/go-rest-api/pkg/domain/todo/repository"
	"github.com/RiverPhillips/go-rest-api/pkg/resources"
	"github.com/go-chi/chi/v5"
)

var _ resources.Handler = (*todoHandlerV1)(nil)

type todoHandlerV1 struct {
	repo repository.Todo
}

// NewTodoHandlerV1 creates a new instance of TodoHandlerV1
func NewTodoHandlerV1(repo repository.Todo) resources.Handler {
	return &todoHandlerV1{
		repo: repo,
	}
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

	id, createdAt, err := t.repo.Create(r.Context(), &req.Data.Attributes)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	created := todo.CreateTodoResponse{
		Data: todo.CreateTodoResponseData{
			Type: "todos",
			Attributes: todo.CreateTodoResponseAttributes{
				Title:       req.Data.Attributes.Title,
				Description: req.Data.Attributes.Description,
				Completed:   false,
				UpdatedAt:   createdAt.Time,
				CreatedAt:   createdAt.Time,
			},
			ID: id,
			Links: resources.Links{
				Self: fmt.Sprintf("http://localhost:8080%s/%d", r.URL.Path, id),
			},
		},
	}

	if err := json.NewEncoder(w).Encode(created); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", created.Data.Links.Self)
	w.WriteHeader(http.StatusCreated)
}
