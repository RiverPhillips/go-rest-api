package resources

// Links is the HATEOAS links
type Links struct {
	Self string `json:"self"`
}

// CreateResponseData represents the data for creating a resource
type CreateResponseData[T any, U any] struct {
	Type       string `json:"type"`
	Attributes T      `json:"attributes"`
	ID         U      `json:"ID"`
	Links      Links  `json:"links"`
}

// CreatedResponse represents the response for creating a resource
type CreatedResponse[T any, U any] struct {
	Data CreateResponseData[T, U] `json:"data"`
}
