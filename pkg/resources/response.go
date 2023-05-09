package resources

type Links struct {
	Self string `json:"self"`
}

type CreateResponseData[T any, U any] struct {
	Type       string `json:"type"`
	Attributes T      `json:"attributes"`
	Id         U      `json:"id"`
	Links      Links  `json:"links"`
}

type CreatedResponse[T any, U any] struct {
	Data CreateResponseData[T, U] `json:"data"`
}
