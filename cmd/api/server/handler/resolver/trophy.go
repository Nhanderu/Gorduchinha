package resolver

import (
	"github.com/Nhanderu/gorduchinha/app/entity"
)

type TrophyResolver struct {
	trophy entity.Trophy
}

func NewTrophyResolver(trophy entity.Trophy) *TrophyResolver {
	return &TrophyResolver{
		trophy: trophy,
	}
}

func (r TrophyResolver) Year() int32 {
	return int32(r.trophy.Year)
}

func (r TrophyResolver) Champ() *ChampResolver {
	return NewChampResolver(r.trophy.Champ)
}
