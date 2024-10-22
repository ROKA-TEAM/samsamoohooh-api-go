package utils

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/repository/database/ent"
	"samsamoohooh-go-api/pkg/box"
)

func Wrap(err error) error {
	var domainErr error
	switch {
	case ent.IsConstraintError(err):
		domainErr = domain.ErrConstraint
	case ent.IsNotFound(err):
		domainErr = domain.ErrNotFound
	case ent.IsNotLoaded(err):
		domainErr = domain.ErrNotLoaded
	case ent.IsNotSingular(err):
		domainErr = domain.ErrNotSingular
	case ent.IsValidationError(err):
		domainErr = domain.ErrValidation
	default:
		return nil
	}
	return box.Wrap(domainErr, err.Error())
}
