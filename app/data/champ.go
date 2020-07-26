package data

import (
	"database/sql"

	"github.com/Nhanderu/gorduchinha/app/contract"
	"github.com/Nhanderu/gorduchinha/app/entity"
	"github.com/pkg/errors"
)

type champRepo struct {
	ex           executor
	entity       string
	selectFields string
}

func newChampRepo(ex executor) contract.ChampRepo {
	return champRepo{
		ex:     ex,
		entity: "champ",
		selectFields: `
			c.id
			, c.created_at
			, c.updated_at
			, c.slug
			, c.name
		`,
	}
}

func (r champRepo) parseEntities(rows *sql.Rows, err error) ([]entity.Champ, error) {
	if err != nil {
		return nil, errors.WithStack(err)
	}

	champs := make([]entity.Champ, 0)
	for rows.Next() {

		champ, err := r.parseEntity(rows)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		champs = append(champs, champ)
	}

	return champs, nil
}

func (r champRepo) parseEntity(s scanner) (entity.Champ, error) {

	var champ entity.Champ
	err := s.Scan(
		&champ.ID,
		&champ.CreatedAt,
		&champ.UpdatedAt,
		&champ.Slug,
		&champ.Name,
	)
	if err != nil {
		return entity.Champ{}, errors.WithStack(err)
	}

	return champ, nil
}

func (r champRepo) FindAll() ([]entity.Champ, error) {
	const query = `
		SELECT %s
			FROM tb_champ AS c
			WHERE c.deleted_at IS NULL
		;
	`

	champs, err := r.parseEntities(r.ex.Query(query))
	if err != nil {
		return nil, errors.WithStack(parseError(err, r.entity))
	}

	return champs, nil
}

func (r champRepo) FindBySlug(slug string) (entity.Champ, error) {
	const query = `
		SELECT %s
			FROM tb_champ AS c
			WHERE c.deleted_at IS NULL
				AND c.slug = $1
		;
	`

	champ, err := r.parseEntity(r.ex.QueryRow(
		query,
		slug,
	))
	if err != nil {
		return entity.Champ{}, errors.WithStack(parseError(err, r.entity))
	}

	return champ, nil
}
