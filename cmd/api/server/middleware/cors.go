package middleware

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

const (
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowHeaders     = "*"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func CORSMiddleware() RequestMiddleware {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {

			if string(ctx.Method()) == http.MethodOptions {
				ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
				ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
				ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
				ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
				ctx.SetStatusCode(http.StatusNoContent)
				return
			}

			next(ctx)
		}
	}
}
