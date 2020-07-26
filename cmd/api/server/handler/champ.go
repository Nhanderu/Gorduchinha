package handler

import (
	"github.com/Nhanderu/gorduchinha/app/contract"
	"github.com/Nhanderu/gorduchinha/cmd/api/server/viewmodel"
	"github.com/valyala/fasthttp"
)

func ListChamps(champService contract.ChampService) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		champs, err := champService.Find()
		if err != nil {
			HandleError(ctx, err)
			return
		}

		respondOK(ctx, viewmodel.ParseChampResponseList(champs))
	}
}

func FindChampBySlug(champService contract.ChampService) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		slug, _ := ctx.UserValue("slug").(string)
		champ, err := champService.FindBySlug(slug)
		if err != nil {
			HandleError(ctx, err)
			return
		}

		respondOK(ctx, viewmodel.ParseChampResponse(champ))
	}
}

func UpdateChamps(scraperService contract.ScraperService) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		err := scraperService.ScrapeAndUpdate()
		if err != nil {
			HandleError(ctx, err)
			return
		}

		respondOK(ctx, nil)
	}
}
