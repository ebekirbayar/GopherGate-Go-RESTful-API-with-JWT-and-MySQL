package model

// BaseResponse represents the structure of a generic response.
type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  bool        `json:"status"`
}

// BaseErrorResponse represents the structure of an error response.
type BaseErrorResponse struct {
	Message string  `json:"message"`
	Errors  []error `json:"errors"`
	Status  bool    `json:"status"`
}
