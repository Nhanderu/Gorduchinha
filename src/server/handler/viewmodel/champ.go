package viewmodel

import (
	"github.com/Nhanderu/gorduchinha/src/domain/entity"
)

type ChampResponse struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

func ParseChampResponseList(teams []entity.Champ) []ChampResponse {

	vm := make([]ChampResponse, len(teams))
	for i := range vm {
		vm[i] = ParseChampResponse(teams[i])
	}

	return vm
}

func ParseChampResponse(champ entity.Champ) ChampResponse {
	return ChampResponse{
		Slug: champ.Slug,
		Name: champ.Name,
	}
}
