package handler

import (
	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	"github.com/Nhanderu/gorduchinha/src/server/handler/viewmodel"
	"github.com/valyala/fasthttp"
)

func ListTeams(teamService contract.TeamService) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		teams, err := teamService.FindAll()
		if err != nil {
			HandleError(ctx, err)
			return
		}

		RespondOK(ctx, viewmodel.ParseTeamResponseList(teams))
	}
}

func FindTeamByAbbr(teamService contract.TeamService) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		abbr, _ := ctx.UserValue("abbr").(string)
		team, err := teamService.FindByAbbr(abbr)
		if err != nil {
			HandleError(ctx, err)
			return
		}

		RespondOK(ctx, viewmodel.ParseTeamResponse(team))
	}
}
