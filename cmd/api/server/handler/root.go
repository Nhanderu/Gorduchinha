package handler

import (
	"github.com/valyala/fasthttp"
)

func ShowAppVersion(version string) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		respondOK(ctx, version)
	}
}

func HealthCheck() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		respondOK(ctx, nil)
	}
}
