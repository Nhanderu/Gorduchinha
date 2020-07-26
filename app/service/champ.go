package service

import (
	"fmt"
	"time"

	"github.com/Nhanderu/gorduchinha/app/contract"
	"github.com/Nhanderu/gorduchinha/app/entity"
	"github.com/pkg/errors"
)

type champService struct {
	data  contract.DataManager
	cache contract.CacheManager
}

func NewChampService(
	data contract.DataManager,
	cache contract.CacheManager,
) contract.ChampService {

	return champService{
		data:  data,
		cache: cache,
	}
}

func (s champService) FindAll() ([]entity.Champ, error) {

	var champs []entity.Champ

	cacheKey := "champ-find-all"
	err := s.cache.GetJSON(cacheKey, &champs)
	if err != nil {

		champs, err = s.data.Champ().FindAll()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		s.cache.SetJSON(cacheKey, champs)
		s.cache.SetExpiration(cacheKey, time.Hour*24*30)
	}

	return champs, nil
}

func (s champService) FindBySlug(slug string) (entity.Champ, error) {

	var champ entity.Champ

	cacheKey := fmt.Sprintf("champ-find-by-slug-%s", slug)
	err := s.cache.GetJSON(cacheKey, &champ)
	if err != nil {

		champ, err = s.data.Champ().FindBySlug(slug)
		if err != nil {
			return champ, errors.WithStack(err)
		}

		s.cache.SetJSON(cacheKey, champ)
		s.cache.SetExpiration(cacheKey, time.Hour*24*30)
	}

	return champ, nil
}
