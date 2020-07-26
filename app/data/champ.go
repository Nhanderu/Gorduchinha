package data

import (
	"github.com/Nhanderu/gorduchinha/app/entity"
	"github.com/pkg/errors"
)

type champRepo struct {
	ex executor
}

func (r champRepo) FindAll() ([]entity.Champ, error) {
	const query = `
		SELECT
			c.id
			, c.slug
			, c.name
			FROM tb_champ AS c
			WHERE c.deleted_at IS NULL
		;
	`

	rows, err := r.ex.Query(query)
	if err != nil {
		return nil, errors.WithStack(parseError(err))
	}
	defer rows.Close()

	champs := make([]entity.Champ, 0)
	for rows.Next() {

		var champ entity.Champ
		err = rows.Scan(
			&champ.ID,
			&champ.Slug,
			&champ.Name,
		)
		if err != nil {
			return nil, errors.WithStack(parseError(err))
		}

		champs = append(champs, champ)
	}

	return champs, nil
}

func (r champRepo) FindBySlug(slug string) (entity.Champ, error) {
	const query = `
		SELECT
			c.id
			, c.slug
			, c.name
			FROM tb_champ AS c
			WHERE c.deleted_at IS NULL
				AND c.slug = $1
		;
	`

	var champ entity.Champ
	err := r.ex.QueryRow(
		query,
		slug,
	).Scan(
		&champ.ID,
		&champ.Slug,
		&champ.Name,
	)
	if err != nil {
		return champ, errors.WithStack(parseError(err))
	}

	return champ, nil
}
