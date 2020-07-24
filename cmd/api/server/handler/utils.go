package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Nhanderu/gorduchinha/app/constant"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

const (
	ErrorCodeContextKey    = "error-code"
	ErrorMessageContextKey = "error-message"
)

type resultWrapper struct {
	Success bool                `json:"success"`
	Data    interface{}         `json:"data,omitempty"`
	Error   *resultWrapperError `json:"error,omitempty"`
}

type resultWrapperError struct {
	Code string `json:"code,omitempty"`
}

func respondOK(ctx *fasthttp.RequestCtx, data interface{}) {
	respondJSON(ctx, http.StatusOK, resultWrapper{
		Success: true,
		Data:    data,
	})
}

func HandleError(ctx *fasthttp.RequestCtx, err error) {

	if err == nil {
		respondInternalError(ctx, "Unknown error.")
		return
	}

	ctx.SetUserValue(ErrorMessageContextKey, err.Error())
	err = errors.Cause(err)

	switch err {

	case constant.ErrNotFound:
		respondNotFoundError(ctx)
		return

	case constant.ErrNotAuthorized:
		respondAuthError(ctx)
		return

	}

	switch e := err.(type) {

	case constant.ValidationError:
		respondValidationError(ctx, e.Error())
		return

	}

	respondInternalError(ctx, "Internal error.")
}

func respondAuthError(ctx *fasthttp.RequestCtx) {
	respondError(ctx, http.StatusUnauthorized, "Access unauthorized.")
}

func respondValidationError(ctx *fasthttp.RequestCtx, code string) {
	respondError(ctx, http.StatusUnprocessableEntity, code)
}

func respondRequestError(ctx *fasthttp.RequestCtx, code string) {
	respondError(ctx, http.StatusBadRequest, code)
}

func respondInternalError(ctx *fasthttp.RequestCtx, code string) {
	respondError(ctx, http.StatusInternalServerError, code)
}

func respondNotFoundError(ctx *fasthttp.RequestCtx) {
	respondError(ctx, http.StatusNotFound, "Resource not found.")
}

func respondError(ctx *fasthttp.RequestCtx, status int, code string) {
	ctx.SetUserValue(ErrorCodeContextKey, code)
	respondJSON(ctx, status, resultWrapper{
		Success: false,
		Error: &resultWrapperError{
			Code: code,
		},
	})
}

func respondJSON(ctx *fasthttp.RequestCtx, code int, result interface{}) {
	respond(ctx, "app/json; charset=UTF-8", code, result)
}

func respond(ctx *fasthttp.RequestCtx, contentType string, code int, result interface{}) {
	ctx.SetContentType(contentType)
	ctx.SetStatusCode(code)
	json.NewEncoder(ctx).Encode(result)
}
