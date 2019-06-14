package service

import (
	"fmt"
	"time"

	"github.com/Nhanderu/gorduchinha/src/domain/entity"
	"github.com/pkg/errors"
)

type champService struct {
	svc *Service
}

func newChampService(svc *Service) champService {
	return champService{
		svc: svc,
	}
}

func (s champService) FindAll() ([]entity.Champ, error) {

	var champs []entity.Champ

	cacheKey := "champ-find-all"
	err := s.svc.cache.GetJSON(cacheKey, &champs)
	if err != nil {

		champs, err = s.svc.db.Champ().FindAll()
		if err != nil {
			return nil, errors.WithStack(err)
		}

		s.svc.cache.SetJSON(cacheKey, champs)
		s.svc.cache.SetExpiration(cacheKey, time.Hour*24*30)
	}

	return champs, nil
}

func (s champService) FindBySlug(slug string) (entity.Champ, error) {

	var champ entity.Champ

	cacheKey := fmt.Sprintf("champ-find-by-slug-%s", slug)
	err := s.svc.cache.GetJSON(cacheKey, &champ)
	if err != nil {

		champ, err = s.svc.db.Champ().FindBySlug(slug)
		if err != nil {
			return champ, errors.WithStack(err)
		}

		s.svc.cache.SetJSON(cacheKey, champ)
		s.svc.cache.SetExpiration(cacheKey, time.Hour*24*30)
	}

	return champ, nil
}
