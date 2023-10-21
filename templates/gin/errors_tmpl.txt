package errors

import "fmt"

type HttpError struct {
	Description string `json:"description,omitempty"`
	StatusCode  int    `json:"status_code,omitempty"`
}

func (err HttpError) Error() string {
	return fmt.Sprintf("descritption: %s", err.Description)
}

func NewHttpError(descritption string, statusCode int) *HttpError {
	return &HttpError{
		Description: descritption,
		StatusCode:  statusCode,
	}
}
