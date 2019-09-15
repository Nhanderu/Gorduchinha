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

func (r TeamResolver) Trophies() []*TrophyResolver {

	resolvers := make([]*TrophyResolver, len(r.team.Trophies))
	for i := range r.team.Trophies {
		resolvers[i] = NewTrophyResolver(r.team.Trophies[i])
	}

	return resolvers
}
