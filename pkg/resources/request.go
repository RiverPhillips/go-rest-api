package resources

type CreateRequest[T any] struct {
	Data CreateRequestData[T] `json:"data"`
}

type CreateRequestData[T any] struct {
	Type       string `json:"type"`
	Attributes T      `json:"attributes"`
	Id         string `json:"id"`
}
