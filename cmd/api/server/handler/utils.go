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

func RespondGraphQL(ctx *fasthttp.RequestCtx, data interface{}) {
	respondJSON(ctx, http.StatusOK, data)
}

func RespondOK(ctx *fasthttp.RequestCtx, data interface{}) {
	respondJSON(ctx, http.StatusOK, resultWrapper{
		Success: true,
		Data:    data,
	})
}

func RespondError(ctx *fasthttp.RequestCtx, status int, code string) {
	ctx.SetUserValue(ErrorCodeContextKey, code)
	respondJSON(ctx, status, resultWrapper{
		Success: false,
		Error: &resultWrapperError{
			Code: code,
		},
	})
}

func RespondNoContent(ctx *fasthttp.RequestCtx) {
	respondJSON(ctx, http.StatusNoContent, resultWrapper{
		Success: true,
	})
}

func RespondAuthError(ctx *fasthttp.RequestCtx) {
	RespondError(ctx, http.StatusUnauthorized, "Access unauthorized.")
}

func RespondValidationError(ctx *fasthttp.RequestCtx, code string) {
	RespondError(ctx, http.StatusUnprocessableEntity, code)
}

func RespondRequestError(ctx *fasthttp.RequestCtx, code string) {
	RespondError(ctx, http.StatusBadRequest, code)
}

func RespondInternalError(ctx *fasthttp.RequestCtx, code string) {
	RespondError(ctx, http.StatusInternalServerError, code)
}

func RespondNotFoundError(ctx *fasthttp.RequestCtx) {
	RespondError(ctx, http.StatusNotFound, "Resource not found.")
}

func HandleError(ctx *fasthttp.RequestCtx, err error) {

	if err == nil {
		RespondInternalError(ctx, "Unknown error.")
		return
	}

	ctx.SetUserValue(ErrorMessageContextKey, err.Error())
	err = errors.Cause(err)

	switch err {

	case constant.ErrNotFound:
		RespondNotFoundError(ctx)
		return

	case constant.ErrNotAuthorized:
		RespondAuthError(ctx)
		return

	}

	switch e := err.(type) {

	case constant.ValidationError:
		RespondValidationError(ctx, e.Error())
		return

	}

	RespondInternalError(ctx, "Internal error.")
}

func respondJSON(ctx *fasthttp.RequestCtx, code int, result interface{}) {

	ctx.SetContentType("app/json; charset=UTF-8")
	ctx.SetStatusCode(code)
	json.NewEncoder(ctx).Encode(result)
}