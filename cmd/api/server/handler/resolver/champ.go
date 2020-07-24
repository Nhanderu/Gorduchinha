package resolver

import (
	"github.com/Nhanderu/gorduchinha/app/entity"
)

type ChampResolver struct {
	champ entity.Champ
}

func NewChampResolver(champ entity.Champ) *ChampResolver {
	return &ChampResolver{
		champ: champ,
	}
}

func (r ChampResolver) Slug() string {
	return r.champ.Slug
}

func (r ChampResolver) Name() string {
	return r.champ.Name
}
