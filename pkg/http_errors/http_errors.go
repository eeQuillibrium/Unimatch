package httpErrors

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	ErrBadGateway          = "Bad Gateway"
	ErrBadRequest          = "Bad request"
	ErrNotFound            = "Not Found"
	ErrInternalServerError = "Internal Server Error"
	ErrUnauthorized        = "Unauthorized"
	ErrRequestTimeout      = "Request Timeout"
)

type RESTError interface {
	Status() int
	Error() string
	ParseErrors()
}

type RestError struct {
	ErrStatus int    `json:"status,omitempty"`
	ErrError  string `json:"error,omitempty"`
}

func NewRestError(errStatus int, errError string) *RestError {
	return &RestError{
		ErrStatus: errStatus,
		ErrError:  errError,
	}
}

func (r *RestError) Status() int {
	return r.ErrStatus
}

func (r *RestError) Error() string {
	return fmt.Sprintf("error: %s, status: %d", r.ErrError, r.ErrStatus)
}

func ParseErrors(err error) *RestError {
	switch {
	case strings.Contains(strings.ToLower(err.Error()), "grpc"):
		return NewRestError(http.StatusBadGateway, ErrBadGateway)

	}
	return NewRestError(http.StatusInternalServerError, ErrInternalServerError)
}
