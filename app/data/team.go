package data

import (
	"github.com/Nhanderu/gorduchinha/app/entity"
	"github.com/pkg/errors"
)

type teamRepo struct {
	ex executor
}

func (r teamRepo) FindAll() ([]entity.Team, error) {
	const query = `
		SELECT
			t.id
			, t.abbr
			, t.name
			, t.full_name
			FROM tb_team AS t
			WHERE t.deleted_at IS NULL
		;
	`

	teams := make([]entity.Team, 0)
	rows, err := r.ex.Query(query)
	if err != nil {
		return nil, errors.WithStack(parseError(err))
	}
	defer rows.Close()

	for rows.Next() {

		var team entity.Team
		err = rows.Scan(
			&team.ID,
			&team.Abbr,
			&team.Name,
			&team.FullName,
		)
		if err != nil {
			return nil, errors.WithStack(parseError(err))
		}

		teams = append(teams, team)
	}

	return teams, nil
}

func (r teamRepo) FindByAbbr(abbr string) (entity.Team, error) {
	const query = `
		SELECT
			t.id
			, t.abbr
			, t.name
			, t.full_name
			FROM tb_team AS t
			WHERE t.deleted_at IS NULL
				AND t.abbr = $1
		;
	`

	var team entity.Team
	err := r.ex.QueryRow(query,
		abbr,
	).Scan(
		&team.ID,
		&team.Abbr,
		&team.Name,
		&team.FullName,
	)
	if err != nil {
		return team, errors.WithStack(parseError(err))
	}

	return team, nil
}
