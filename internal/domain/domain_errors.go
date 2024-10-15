package domain

import "github.com/pkg/errors"

var (
	ErrTokenGenerate      = errors.New("token generate error")
	ErrTokenParse         = errors.New("token parse error")
	ErrInvalidTokenIssuer = errors.New("invalid token issuer")
	ErrTokenNotActiveYet  = errors.New("token is not active yet")
	ErrTokenExpired       = errors.New("token is expired")
)
