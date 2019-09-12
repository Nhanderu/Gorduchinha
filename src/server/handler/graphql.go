package handler

import (
	"io/ioutil"

	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	gqltools "github.com/bhoriuchi/graphql-go-tools"
	"github.com/graphql-go/graphql"
	"github.com/valyala/fasthttp"
)

func HandleGraphql(
	teamService contract.TeamService,
	champService contract.ChampService,
	scraperService contract.ScraperService,
) func(ctx *fasthttp.RequestCtx) {

	typedefs, _ := ioutil.ReadFile("static/graphql/schema.gql")
	schema, _ := gqltools.MakeExecutableSchema(gqltools.ExecutableSchema{
		TypeDefs: typedefs,
		Resolvers: map[string]interface{}{
			"Team": &gqltools.ObjectResolver{
				Fields: gqltools.FieldResolveMap{
					"name": func(p graphql.ResolveParams) (interface{}, error) {
						return "name", nil
					},
				},
			},
			"Champ": &gqltools.ObjectResolver{
				Fields: gqltools.FieldResolveMap{
					"name": func(p graphql.ResolveParams) (interface{}, error) {
						return "name", nil
					},
				},
			},
		},
	})

	return func(ctx *fasthttp.RequestCtx) {
		RespondOK(ctx, graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: string(ctx.PostBody()),
		}))
	}
}
