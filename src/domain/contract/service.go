package contract

import (
	"github.com/Nhanderu/gorduchinha/src/domain/entity"
)

type TeamService interface {
	FindAll() ([]entity.Team, error)
	FindByAbbr(abbr string) (entity.Team, error)
}

type ChampService interface {
	FindAll() ([]entity.Champ, error)
	FindBySlug(slug string) (entity.Champ, error)
}

type ScraperService interface {
	ScrapeAndUpdate() error
}
