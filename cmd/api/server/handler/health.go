package handler

import (
	"github.com/valyala/fasthttp"
)

func HealthCheck() func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		respondOK(ctx, nil)
	}
}
