package utils

import (
	"samsamoohooh-go-api/internal/infra/exception"
	"samsamoohooh-go-api/internal/infra/storage/mysql/ent"
)

func Wrap(err error) error {
	switch {
	case err == nil:
		return err

	case ent.IsConstraintError(err):
		return exception.New(
			err,
			exception.ErrMySQLConstraint,
			exception.StatusBadRequest,
			"database constraint error",
		)

	case ent.IsNotFound(err):
		return exception.New(
			err,
			exception.ErrMySQLNotFound,
			exception.StatusNotFound,
			"database not found error",
		)

	case ent.IsNotLoaded(err):
		return exception.New(
			err,
			exception.ErrMySQLNotLoaded,
			exception.StatusInternalServerError,
			"database not loaded error",
		)

	case ent.IsNotSingular(err):
		return exception.New(
			err,
			exception.ErrMySQLNotSingular,
			exception.StatusInternalServerError,
			"database not singular error",
		)

	case ent.IsValidationError(err):
		return exception.New(
			err,
			exception.ErrMySQLValidation,
			exception.StatusBadRequest,
			"database validation error",
		)
	}

	return exception.New(
		err,
		exception.ErrMysqlInternal,
		exception.StatusInternalServerError,
		"database internal error",
	)
}
