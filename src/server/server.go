package server

import (
	"fmt"
	"net/http"

	"github.com/Nhanderu/gorduchinha/src/domain/service"
	"github.com/Nhanderu/gorduchinha/src/infra/config"
	"github.com/Nhanderu/gorduchinha/src/infra/logger"
	"github.com/Nhanderu/gorduchinha/src/server/handler"
	"github.com/Nhanderu/gorduchinha/src/server/middleware"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

// Run setup and executes the server
func Run(svc *service.Service, cfg config.Config, log logger.Logger) error {

	address := fmt.Sprintf(":%d", cfg.Server.Port)
	router := registerRoutes(svc, cfg, log)

	err := fasthttp.ListenAndServe(address, router)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// registerRoutes registers the API routes.
func registerRoutes(svc *service.Service, cfg config.Config, log logger.Logger) fasthttp.RequestHandler {

	router := newRouter()

	open := router.group(cfg.Server.Prefix, middleware.LoggerMiddleware(log))
	open.handle(http.MethodGet, "/health", handler.HealthCheck())
	open.handle(http.MethodGet, "/version", handler.ShowAppVersion(cfg.App.Version))
	open.handle(http.MethodGet, "/champs", handler.ListChamps(svc))
	open.handle(http.MethodGet, "/champs/:slug", handler.FindChampBySlug(svc))
	open.handle(http.MethodGet, "/teams", handler.ListTeams(svc))
	open.handle(http.MethodGet, "/teams/:abbr", handler.FindTeamByAbbr(svc))

	return router.requestHandler()
}
