package middleware

import (
	"github.com/Nhanderu/gorduchinha/app/constant"
	"github.com/Nhanderu/gorduchinha/cmd/api/server/handler"
	"github.com/valyala/fasthttp"
)

func QueryKeyValidation(queryKey string) RequestMiddleware {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {

			if string(ctx.QueryArgs().Peek("key")) != queryKey {
				handler.HandleError(ctx, constant.NewErrorInvalidQueryKey())
				return
			}

			next(ctx)
		}
	}
}
