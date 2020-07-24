package data

import (
	"database/sql"

	"github.com/Nhanderu/gorduchinha/app/constant"
	"github.com/pkg/errors"
)

func parseError(err error) error {
	if err == nil {
		return nil
	}

	originalErr := errors.Cause(err)

	switch originalErr {
	case sql.ErrNoRows:
		return errors.WithStack(constant.ErrNotFound)
	}

	return errors.WithStack(err)
}
