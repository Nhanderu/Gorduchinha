package handler

import (
	"github.com/valyala/fasthttp"
)

func HealthCheck() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		respondOK(ctx, nil)
	}
}
