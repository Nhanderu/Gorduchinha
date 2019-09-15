package resolver

import (
	"github.com/Nhanderu/gorduchinha/src/domain/contract"
)

type QueryResolver struct {
	teamService  contract.TeamService
	champService contract.ChampService
}

func NewQueryResolver(teamService contract.TeamService, champService contract.ChampService) *QueryResolver {
	return &QueryResolver{
		teamService:  teamService,
		champService: champService,
	}
}

type TeamArgs struct {
	Abbr string
}

func (r QueryResolver) Team(args *TeamArgs) *TeamResolver {

	team, _ := r.teamService.FindByAbbr(args.Abbr)
	return NewTeamResolver(team)
}

func (r QueryResolver) Teams() []*TeamResolver {
	teams, _ := r.teamService.FindAll()
	resolvers := make([]*TeamResolver, len(teams))
	for i := range teams {
		resolvers[i] = &TeamResolver{
			team: teams[i],
		}
	}

	return resolvers
}

type ChampArgs struct {
	Slug string
}

func (r QueryResolver) Champ(args *ChampArgs) *ChampResolver {

	champ, _ := r.champService.FindBySlug(args.Slug)
	return NewChampResolver(champ)
}

func (r QueryResolver) Champs() []*ChampResolver {
	champs, _ := r.champService.FindAll()
	resolvers := make([]*ChampResolver, len(champs))
	for i := range champs {
		resolvers[i] = &ChampResolver{
			champ: champs[i],
		}
	}

	return resolvers
}
