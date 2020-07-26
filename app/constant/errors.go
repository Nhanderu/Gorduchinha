package constant

import "fmt"

const (

	// ErrorCodeNotFound means that a resource was not found.
	ErrorCodeNotFound = "not-found"

	// ErrorCodeCacheMiss indicates a cache miss when fetching an item from CacheManager.
	ErrorCodeCacheMiss = "cache-miss"

	// ErrorCodeTooManyRequests means that an IP is sending too much requests and we're
	// blocking it.
	ErrorCodeTooManyRequests = "too-many-requests"

	// ErrorCodeInvalidRequestBody means that the HTTP request body has an invalid
	// format.
	ErrorCodeInvalidRequestBody = "invalid-request-body"

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

func NewErrorNotFound(field string) AppError {
	return AppError{
		Field: field,
		Code:  ErrorCodeNotFound,
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

func NewErrorInternal() AppError {
	return AppError{
		Code: ErrorCodeInternal,
	}
}
