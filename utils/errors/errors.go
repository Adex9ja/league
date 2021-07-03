package errors

import (
	"errors"
	"net/http"
)

type ErrorResponser struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) *ErrorResponser {
	return &ErrorResponser{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *ErrorResponser {
	return &ErrorResponser{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *ErrorResponser {
	return &ErrorResponser{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
