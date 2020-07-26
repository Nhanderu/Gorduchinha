package middleware

import (
	"strconv"

	"github.com/Nhanderu/gorduchinha/app/constant"
	"github.com/Nhanderu/gorduchinha/cmd/api/server/handler"
	"github.com/valyala/fasthttp"
)

func BodyLimit() RequestMiddleware {

	const (
		limit = 32 * 1024
	)

	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {

			contentLength, _ := strconv.Atoi(string(ctx.Request.Header.Peek("Content-Length")))
			if contentLength > limit {
				handler.HandleError(ctx, constant.NewErrorRequestBodyTooLarge())
				return
			}

			next(ctx)
		}
	}
}
