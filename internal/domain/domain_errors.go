package domain

import "github.com/pkg/errors"

var (
	ErrInvalidTokenIssuer = errors.New("invalid token issuer")
	ErrTokenExpired       = errors.New("token is expired")
	ErrTokenNotActiveYet  = errors.New("token is not active yet")
	ErrTokenParse         = errors.New("token parse error")
	ErrTokenGenerate      = errors.New("token generate error")
)
