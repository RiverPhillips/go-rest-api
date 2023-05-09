package resources

import "github.com/go-chi/chi/v5"

type Handler interface {
	Route() func(r chi.Router)
}
