package server

import (
	"github.com/Nhanderu/gorduchinha/src/server/middleware"
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
	root.router.Handle(method, root.prefix+path, middleware.Use(handler, root.mw...))
}
