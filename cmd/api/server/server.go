package server

import (
	"fmt"
	"net/http"

	"github.com/Nhanderu/gorduchinha/app/contract"
	"github.com/Nhanderu/gorduchinha/app/logger"
	"github.com/Nhanderu/gorduchinha/cmd/api/server/handler"
	"github.com/Nhanderu/gorduchinha/cmd/api/server/middleware"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

func Run(
	serverPort int,
	serverPrefix string,
	log logger.Logger,
	teamService contract.TeamService,
	champService contract.ChampService,
	scraperService contract.ScraperService,
) error {

	router := newRouter()

	open := router.group(
		serverPrefix,
		middleware.LoggerMiddleware(log),
		middleware.CORSMiddleware(),
	)

	open.handle(http.MethodGet, "/health", handler.HealthCheck())
	open.handle(http.MethodPost, "/graphql", handler.HandleGraphql(teamService, champService))
	open.handle(http.MethodGet, "/teams", handler.ListTeams(teamService))
	open.handle(http.MethodGet, "/teams/{abbr}", handler.FindTeamByAbbr(teamService))
	open.handle(http.MethodGet, "/champs", handler.ListChamps(champService))
	open.handle(http.MethodPut, "/champs", handler.UpdateChamps(scraperService))
	open.handle(http.MethodGet, "/champs/{slug}", handler.FindChampBySlug(champService))

	address := fmt.Sprintf(":%d", serverPort)
	err := fasthttp.ListenAndServe(address, router.requestHandler())
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
