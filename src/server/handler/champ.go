package handler

import (
	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	"github.com/Nhanderu/gorduchinha/src/server/handler/viewmodel"
	"github.com/valyala/fasthttp"
)

func ListChamps(champService contract.ChampService) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		champs, err := champService.FindAll()
		if err != nil {
			HandleError(ctx, err)
			return
		}

		RespondOK(ctx, viewmodel.ParseChampResponseList(champs))
	}
}

func FindChampBySlug(champService contract.ChampService) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		slug, _ := ctx.UserValue("slug").(string)
		champ, err := champService.FindBySlug(slug)
		if err != nil {
			HandleError(ctx, err)
			return
		}

		RespondOK(ctx, viewmodel.ParseChampResponse(champ))
	}
}
