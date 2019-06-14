package middleware

import (
	"github.com/valyala/fasthttp"
)

func Use(h fasthttp.RequestHandler, mws ...RequestMiddleware) fasthttp.RequestHandler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

type RequestMiddleware func(fasthttp.RequestHandler) fasthttp.RequestHandler
