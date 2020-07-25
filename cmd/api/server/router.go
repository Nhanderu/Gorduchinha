package server

import (
	"net/http"

	"github.com/Nhanderu/gorduchinha/cmd/api/server/middleware"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type r struct {
	router *router.Router
	prefix string
	mw     []middleware.RequestMiddleware
}

func newRouter() *r {
	return &r{
		router: router.New(),
		mw:     make([]middleware.RequestMiddleware, 0),
	}
}

func (root *r) requestHandler() fasthttp.RequestHandler {
	return root.router.Handler
}

func (root *r) group(prefix string, mws ...middleware.RequestMiddleware) *r {
	return &r{
		router: root.router,
		prefix: root.prefix + prefix,
		mw:     append(root.mw, mws...),
	}
}

func (root *r) handle(method, path string, handler fasthttp.RequestHandler) {
	p := root.prefix + path
	root.router.Handle(method, p, middleware.Use(handler, root.mw...))

	_, registered := preflightRegistered[p]
	if !registered {
		root.router.Handle(http.MethodOptions, p, middleware.Use(preflightHandler, root.mw...))
		preflightRegistered[p] = struct{}{}
	}
}

var (
	preflightRegistered = map[string]struct{}{}
)

func preflightHandler(ctx *fasthttp.RequestCtx) {}
