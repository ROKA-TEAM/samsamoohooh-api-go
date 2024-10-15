package utils

import (
	"github.com/pkg/errors"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/repository/database/ent"
)

func Wrap(err error) error {
	var domainErr error
	switch {

	case ent.IsNotFound(err):
		domainErr = domain.ErrNotFound
	case ent.IsNotLoaded(err):
		domainErr = domain.ErrNotLoaded
	case ent.IsConstraintError(err):
		domainErr = domain.ErrConstraint
	case ent.IsNotSingular(err):
		domainErr = domain.ErrNotSingular
	case ent.IsValidationError(err):
		domainErr = domain.ErrValidation
	}

	return errors.Wrap(domainErr, err.Error())
}
