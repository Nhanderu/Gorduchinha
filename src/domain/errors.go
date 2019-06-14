package domain

import (
	"fmt"
	"net/http"
)

type errorString string

func (err errorString) Error() string {
	return string(err)
}

const (
	ErrNotFound      = errorString("resource not found")
	ErrCacheMiss     = errorString("miss from cache")
	ErrNotAuthorized = errorString("access unauthorized")
)

type ValidationError struct {
	Field string
}

var _ error = ValidationError{}

func (e ValidationError) Error() string {
	return fmt.Sprintf(`field "%s" is invalid`, e.Field)
}

type HTTPError struct {
	Status int
}

var _ error = HTTPError{}

func (e HTTPError) Error() string {
	return http.StatusText(e.Status)
}
