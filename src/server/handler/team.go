package handler

import (
	"github.com/Nhanderu/gorduchinha/src/domain/service"
	"github.com/Nhanderu/gorduchinha/src/server/handler/viewmodel"
	"github.com/valyala/fasthttp"
)

func ListTeams(svc *service.Service) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		teams, err := svc.Team.FindAll()
		if err != nil {
			HandleError(ctx, err)
			return
		}

		RespondOK(ctx, viewmodel.ParseTeamResponseList(teams))
	}
}

func FindTeamByAbbr(svc *service.Service) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		abbr, _ := ctx.UserValue("abbr").(string)
		team, err := svc.Team.FindByAbbr(abbr)
		if err != nil {
			HandleError(ctx, err)
			return
		}

		RespondOK(ctx, viewmodel.ParseTeamResponse(team))
	}
}
