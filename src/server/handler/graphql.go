package handler

import (
	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	"github.com/valyala/fasthttp"
)

func HandleGraphql(
	teamService contract.TeamService,
	champService contract.ChampService,
	scraperService contract.ScraperService,
) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		RespondNoContent(ctx)
	}
}
