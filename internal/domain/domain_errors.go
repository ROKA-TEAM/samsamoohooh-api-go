package domain

import "github.com/pkg/errors"

var (
	ErrTokenGenerate              = errors.New("token generate error")
	ErrTokenParse                 = errors.New("token parse error")
	ErrInvalidTokenIssuer         = errors.New("invalid token issuer")
	ErrTokenNotActiveYet          = errors.New("token is not active yet")
	ErrTokenExpired               = errors.New("token is expired")
	ErrMissingAuthorizationHeader = errors.New("missing authorization header")

	ErrNotFound    = errors.New("not found")
	ErrNotLoaded   = errors.New("not loaded")
	ErrConstraint  = errors.New("constraint error")
	ErrNotSingular = errors.New("singular error")
	ErrValidation  = errors.New("validation error")

	ErrInternal      = errors.New("internal error")
	ErrNotMatchState = errors.New("not match state")
)
