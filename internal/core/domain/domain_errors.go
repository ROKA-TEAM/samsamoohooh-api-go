package domain

import "github.com/pkg/errors"

var (
	// ErrBadParam = errors.New("given param is not valid") // 400 Bad Request // 501 Not Implemented“
	ErrUnauthorized = errors.New("unauthorized")
	ErrInternal     = errors.New("internal error")
)
