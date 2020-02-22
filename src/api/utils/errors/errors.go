package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type APiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	AStatus  int    `json:"status"`
	AMessage string `json:"message"`
	AnError  string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.AStatus
}

func (e *apiError) Message() string {
	return e.AMessage
}

func (e *apiError) Error() string {
	return e.AnError
}

func NewApiError(statusCode int, message string) APiError {
	return &apiError{
		AStatus:  statusCode,
		AMessage: message,
	}
}

func NewAPiErrorFoBytes(body []byte) (APiError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("Invalid json body")
	}
	return &result, nil
}

func NewInternalServerError(message string) APiError {
	return &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
	}
}

func NewNotFoundApiError(message string) APiError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}

func NewBadRequestError(message string) APiError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
