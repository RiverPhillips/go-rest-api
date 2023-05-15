// Package resources represents the REST API resources
package resources

// CreateRequest represents the request for creating a resource
type CreateRequest[T any] struct {
	Data CreateRequestData[T] `json:"data"`
}

// CreateRequestData represents the data for creating a resource
type CreateRequestData[T any] struct {
	Type       string `json:"type"`
	Attributes T      `json:"attributes"`
	ID         string `json:"ID"`
}
