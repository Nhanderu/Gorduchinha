package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nhanderu/gorduchinha/app/constant"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

const (
	ErrorCodeContextKey    = "error-code"
	ErrorMessageContextKey = "error-message"
)

var (
	errorStatusMap = map[string]int{
		constant.ErrorCodePageNotFound:    http.StatusNotFound,
		constant.ErrorCodeEntityNotFound:  http.StatusNotFound,
		constant.ErrorCodeCacheMiss:       http.StatusInternalServerError,
		constant.ErrorCodeTooManyRequests: http.StatusTooManyRequests,
		constant.ErrorCodeInternal:        http.StatusInternalServerError,
	}
)

type resultWrapper struct {
	Success bool                 `json:"success"`
	Data    interface{}          `json:"data,omitempty"`
	Errors  []resultWrapperError `json:"errors,omitempty"`
}

type resultWrapperError struct {
	Code  string `json:"code"`
	Field string `json:"field,omitempty"`
}

func HandleError(ctx *fasthttp.RequestCtx, err error) {

	if err == nil {
		return
	}

	ctx.SetUserValue(ErrorMessageContextKey, err.Error())
	err = errors.Cause(err)

	switch e := err.(type) {

	case constant.AppError:
		statusCode, ok := errorStatusMap[e.Code]
		if !ok {
			statusCode = http.StatusInternalServerError
		}

		respondError(ctx, statusCode, e.Code, e.Field, e.Error())
		return

	default:
		respondError(
			ctx,
			http.StatusInternalServerError,
			constant.ErrorCodeInternal,
			"",
			fmt.Sprintf("unmapped error: %s", e.Error()),
		)
		return
	}

}

func respondError(ctx *fasthttp.RequestCtx, status int, code string, field string, message string) {
	ctx.SetUserValue(ErrorCodeContextKey, code)
	ctx.SetUserValue(ErrorMessageContextKey, message)

	errors := make([]resultWrapperError, 0)
	errors = append(errors, resultWrapperError{
		Code:  code,
		Field: field,
	})

	respondJSON(ctx, status, resultWrapper{
		Success: false,
		Errors:  errors,
	})
}

func respondOK(ctx *fasthttp.RequestCtx, data interface{}) {
	respondJSON(ctx, http.StatusOK, resultWrapper{
		Success: true,
		Data:    data,
	})
}

func respondJSON(ctx *fasthttp.RequestCtx, code int, result interface{}) {
	ctx.SetContentType("app/json; charset=UTF-8")
	ctx.SetStatusCode(code)
	json.NewEncoder(ctx).Encode(result)
}
