package handler

import (
	"github.com/Nhanderu/gorduchinha/src/domain/service"
	"github.com/Nhanderu/gorduchinha/src/server/handler/viewmodel"
	"github.com/valyala/fasthttp"
)

func ListChamps(svc *service.Service) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		champs, err := svc.Champ.FindAll()
		if err != nil {
			HandleError(ctx, err)
			return
		}

		RespondOK(ctx, viewmodel.ParseChampResponseList(champs))
	}
}

func FindChampBySlug(svc *service.Service) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {

		slug, _ := ctx.UserValue("slug").(string)
		champ, err := svc.Champ.FindBySlug(slug)
		if err != nil {
			HandleError(ctx, err)
			return
		}

		RespondOK(ctx, viewmodel.ParseChampResponse(champ))
	}
}
