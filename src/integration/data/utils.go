package data

import (
	"database/sql"

	"github.com/Nhanderu/gorduchinha/src/domain"
	"github.com/pkg/errors"
)

func parseError(err error) error {
	if err == nil {
		return nil
	}

	originalErr := errors.Cause(err)

	switch originalErr {
	case sql.ErrNoRows:
		return errors.WithStack(domain.ErrNotFound)
	}

	return errors.WithStack(err)
}
