package contract

import (
	"github.com/Nhanderu/gorduchinha/src/domain/entity"
)

type DataManager interface {
	RepoManager
	Begin() (TransactionManager, error)
	Close() error
}

type TransactionManager interface {
	RepoManager
	Rollback() error
	Commit() error
}

type RepoManager interface {
	Champ() ChampRepo
	Team() TeamRepo
	Trophy() TrophyRepo
}

type ChampRepo interface {
	FindAll() ([]entity.Champ, error)
	FindBySlug(slug string) (entity.Champ, error)
}

type TeamRepo interface {
	FindAll() ([]entity.Team, error)
	FindByAbbr(abbr string) (entity.Team, error)
}

type TrophyRepo interface {
	FindByTeamID(teamID int) ([]entity.Trophy, error)
	Insert(teamID int, trophy entity.Trophy) error
	DeleteAll() error
}
