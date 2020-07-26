package constant

import "fmt"

const (

	// ErrorCodePageNotFound means that a route was not found.
	ErrorCodePageNotFound = "page-not-found"

	// ErrorCodeMethodNotAllowed means that the current method is not allowed for the
	// specific route.
	ErrorCodeMethodNotAllowed = "method-not-allowed"

	// ErrorCodeCacheMiss indicates a cache miss when fetching an item from CacheManager.
	ErrorCodeCacheMiss = "cache-miss"

	// ErrorCodeTooManyRequests means that an IP is sending too much requests and we're
	// blocking it.
	ErrorCodeTooManyRequests = "too-many-requests"

	// ErrorCodeInvalidRequestBody means that the HTTP request body has an invalid
	// format.
	ErrorCodeInvalidRequestBody = "invalid-request-body"

	// ErrorCodeEntityNotFound means that a resource entity was not found.
	ErrorCodeEntityNotFound = "entity-not-found"

	// ErrorCodeInternal means any general internal error.
	ErrorCodeInternal = "internal"
)

type AppError struct {
	Code  string
	Field string
}

func (e AppError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("error code %s on field %s", e.Code, e.Field)
	}
	return fmt.Sprintf("error code %s", e.Code)
}

func (e AppError) Extensions() map[string]interface{} {

	m := make(map[string]interface{})
	m["code"] = e.Code

	if e.Field != "" {
		m["field"] = e.Field
	}

	return m
}

func NewErrorPageNotFound() AppError {
	return AppError{
		Code: ErrorCodePageNotFound,
	}
}

func NewErrorMethodNotAllowed() AppError {
	return AppError{
		Code: ErrorCodeMethodNotAllowed,
	}
}

func NewErrorCacheMiss() AppError {
	return AppError{
		Code: ErrorCodeCacheMiss,
	}
}

func NewErrorTooManyRequests() AppError {
	return AppError{
		Code: ErrorCodeTooManyRequests,
	}
}

func NewErrorInvalidRequestBody() AppError {
	return AppError{
		Code: ErrorCodeInvalidRequestBody,
	}
}

func NewErrorEntityNotFound(field string) AppError {
	return AppError{
		Field: field,
		Code:  ErrorCodeEntityNotFound,
	}
}

func NewErrorInternal() AppError {
	return AppError{
		Code: ErrorCodeInternal,
	}
}
