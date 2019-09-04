package handler

import (
	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	"github.com/valyala/fasthttp"
)

func RunScraper(scraperService contract.ScraperService) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		err := scraperService.ScrapeAndUpdate()
		if err != nil {
			HandleError(ctx, err)
			return
		}

		RespondNoContent(ctx)
	}
}
