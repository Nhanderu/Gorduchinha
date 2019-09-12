package handler

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	"github.com/Nhanderu/gorduchinha/src/server/handler/viewmodel"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/valyala/fasthttp"
)

type Query struct {
	team  Team
	champ Champ
}

func (q Query) Team() Team {
	return q.team
}

func (q Query) Champ() Champ {
	return q.champ
}

type Team struct{}

func (Team) Name() string {
	return "name"
}

type Champ struct{}

func (Champ) Name() string {
	return "name"
}

func HandleGraphql(
	teamService contract.TeamService,
	champService contract.ChampService,
	scraperService contract.ScraperService,
) func(ctx *fasthttp.RequestCtx) {

	typedefs, _ := ioutil.ReadFile("static/graphql/schema.gql")
	schema := graphql.MustParseSchema(string(typedefs), &Query{})

	return func(ctx *fasthttp.RequestCtx) {

		var request viewmodel.GraphQLQueryRequest
		json.Unmarshal(ctx.PostBody(), &request)

		RespondOK(ctx, schema.Exec(ctx, request.Query, request.OperationName, request.Variables))
	}
}
