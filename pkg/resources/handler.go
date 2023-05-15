package resources

import "github.com/go-chi/chi/v5"

// Handler is the interface for http handlers
type Handler interface {
	Route() func(r chi.Router)
}
