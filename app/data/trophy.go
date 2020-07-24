package data

import (
	"github.com/Nhanderu/gorduchinha/app/entity"
	"github.com/pkg/errors"
)

type trophyRepo struct {
	ex executor
}

func (r trophyRepo) FindByTeamID(teamID int) ([]entity.Trophy, error) {
	const query = `
		SELECT
			t.id
			, t.year
			, c.id
			, c.slug
			, c.name
			FROM tb_trophy AS t
			JOIN tb_champ AS c
				ON t.champ_id = c.id
				AND c.deleted_at IS NULL
			WHERE t.deleted_at IS NULL
				AND t.team_id = $1
		;
	`

	trophies := make([]entity.Trophy, 0)
	rows, err := r.ex.Query(query,
		teamID,
	)
	if err != nil {
		return nil, errors.WithStack(parseError(err))
	}
	defer rows.Close()

	for rows.Next() {

		var trophy entity.Trophy
		err = rows.Scan(
			&trophy.ID,
			&trophy.Year,
			&trophy.Champ.ID,
			&trophy.Champ.Slug,
			&trophy.Champ.Name,
		)
		if err != nil {
			return nil, errors.WithStack(parseError(err))
		}

		trophies = append(trophies, trophy)
	}

	return trophies, nil
}

func (r trophyRepo) Insert(teamID int, trophy entity.Trophy) error {
	const query = `
		INSERT INTO tb_trophy
			( uuid
			, year
			, champ_id
			, team_id
			)
			VALUES
				( UUID_GENERATE_V1()
				, $1
				, $2
				, $3
				)
		;
	`

	_, err := r.ex.Exec(query,
		trophy.Year,
		trophy.Champ.ID,
		teamID,
	)
	if err != nil {
		return errors.WithStack(parseError(err))
	}

	return nil
}

func (r trophyRepo) DeleteAll() error {
	const query = `
		UPDATE tb_trophy
			SET deleted_at = NOW()
			WHERE deleted_at IS NULL
		;
	`

	_, err := r.ex.Exec(query)
	if err != nil {
		return errors.WithStack(parseError(err))
	}

	return nil
}
