package resolver

import (
	"github.com/Nhanderu/gorduchinha/src/domain/entity"
)

type TeamResolver struct {
	team entity.Team
}

func NewTeamResolver(team entity.Team) *TeamResolver {
	return &TeamResolver{
		team: team,
	}
}

func (r TeamResolver) Abbr() string {
	return r.team.Abbr
}

func (r TeamResolver) Name() string {
	return r.team.Name
}

func (r TeamResolver) FullName() string {
	return r.team.FullName
}

func (r TeamResolver) Trophies(args *TrophyArgs) []*TrophyResolver {

	resolvers := make([]*TrophyResolver, len(r.team.Trophies))
	for i, trophy := range r.team.Trophies {
		if args.ChampSlug == nil || *args.ChampSlug == trophy.Champ.Slug {
			resolvers[i] = NewTrophyResolver(trophy)
		}
	}

	return resolvers
}

type TrophyArgs struct {
	ChampSlug *string
}
