package resolver

import (
	"github.com/Nhanderu/gorduchinha/src/domain/contract"
	"github.com/Nhanderu/gorduchinha/src/domain/entity"
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

func (r QueryResolver) Team(args *TeamArgs) *TeamResolver {

	team, _ := r.teamService.FindByAbbr(args.Abbr)
	return &TeamResolver{
		team: team,
	}
}

func (r QueryResolver) Teams() *[]*TeamResolver {
	teams, _ := r.teamService.FindAll()
	resolvers := make([]*TeamResolver, len(teams))
	for i := range teams {
		resolvers[i] = &TeamResolver{
			team: teams[i],
		}
	}

	return &resolvers
}

func (r QueryResolver) Champ(args *ChampArgs) *ChampResolver {

	champ, _ := r.champService.FindBySlug(args.Slug)
	return &ChampResolver{
		champ: champ,
	}
}

func (r QueryResolver) Champs() *[]*ChampResolver {
	champs, _ := r.champService.FindAll()
	resolvers := make([]*ChampResolver, len(champs))
	for i := range champs {
		resolvers[i] = &ChampResolver{
			champ: champs[i],
		}
	}

	return &resolvers
}

type TeamArgs struct {
	Abbr string
}

type TeamResolver struct {
	team entity.Team
}

func (r TeamResolver) Name() string {
	return r.team.Name
}

type ChampArgs struct {
	Slug string
}

type ChampResolver struct {
	champ entity.Champ
}

func (r ChampResolver) Name() string {
	return r.champ.Name
}
