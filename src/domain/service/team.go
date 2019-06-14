package service

import (
	"fmt"
	"time"

	"github.com/Nhanderu/gorduchinha/src/domain/entity"
	"github.com/pkg/errors"
)

type teamService struct {
	svc *Service
}

func newTeamService(svc *Service) teamService {
	return teamService{
		svc: svc,
	}
}

func (s teamService) FindAll() ([]entity.Team, error) {

	var teams []entity.Team

	cacheKey := "team-find-all"
	err := s.svc.cache.GetJSON(cacheKey, &teams)
	if err != nil {

		teams, err = s.svc.db.Team().FindAll()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		for i := range teams {
			teams[i].Trophies, err = s.svc.db.Trophy().FindByTeamID(teams[i].ID)
			if err != nil {
				return nil, errors.WithStack(err)
			}
		}

		s.svc.cache.SetJSON(cacheKey, teams)
		s.svc.cache.SetExpiration(cacheKey, time.Hour*24*30)
	}

	return teams, nil
}

func (s teamService) FindByAbbr(abbr string) (entity.Team, error) {

	var team entity.Team

	cacheKey := fmt.Sprintf("team-find-by-abbr-%s", abbr)
	err := s.svc.cache.GetJSON(cacheKey, &team)
	if err != nil {

		team, err = s.svc.db.Team().FindByAbbr(abbr)
		if err != nil {
			return team, errors.WithStack(err)
		}

		team.Trophies, err = s.svc.db.Trophy().FindByTeamID(team.ID)
		if err != nil {
			return team, errors.WithStack(err)
		}

		s.svc.cache.SetJSON(cacheKey, team)
		s.svc.cache.SetExpiration(cacheKey, time.Hour*24*30)
	}

	return team, nil
}
