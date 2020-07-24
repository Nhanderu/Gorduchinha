package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Nhanderu/gorduchinha/app/contract"
	"github.com/Nhanderu/gorduchinha/cmd/api/server/handler/resolver"
	"github.com/Nhanderu/gorduchinha/cmd/api/server/handler/viewmodel"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/valyala/fasthttp"
)

func HandleGraphql(teamService contract.TeamService, champService contract.ChampService) func(ctx *fasthttp.RequestCtx) {

	typedefs, _ := ioutil.ReadFile("static/graphql/schema.gql")
	queryResolver := resolver.NewQueryResolver(teamService, champService)
	schema := graphql.MustParseSchema(string(typedefs), queryResolver)

	return func(ctx *fasthttp.RequestCtx) {

		var request viewmodel.GraphQLQueryRequest
		err := json.Unmarshal(ctx.PostBody(), &request)
		if err != nil {
			respondRequestError(ctx, "invalid body")
			return
		}

		respond(
			ctx,
			"application/graphql",
			http.StatusOK,
			schema.Exec(ctx, request.Query, request.OperationName, request.Variables),
		)
	}
}
