package utils

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/pkg/box"
	"samsamoohooh-go-api/pkg/token"
)

func Wrap(err error) error {
	var domainErr error
	switch {
	case box.Equal(err, token.ErrTokenParse):
	case box.Equal(err, token.ErrInvalidTokenIssuer):
	case box.Equal(err, token.ErrTokenExpired):
	case box.Equal(err, token.ErrTokenNotActiveYet):
		domainErr = domain.ErrAuthorization
	case err == nil:
		return nil
	default:
		domainErr = domain.ErrInternal
	}

	return box.Wrap(domainErr, err.Error())
}
