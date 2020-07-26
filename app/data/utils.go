package data

import (
	"database/sql"

	"github.com/Nhanderu/gorduchinha/app/constant"
	"github.com/pkg/errors"
)

type executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func parseError(err error, entity string) error {
	if err == nil {
		return nil
	}

	originalErr := errors.Cause(err)

	switch originalErr {
	case sql.ErrNoRows:
		return errors.WithStack(constant.NewErrorNotFound(entity))
	}

	return errors.WithStack(err)
}
