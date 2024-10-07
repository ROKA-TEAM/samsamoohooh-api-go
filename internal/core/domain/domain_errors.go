package domain

import "errors"

var (
	ErrBadParam = errors.New("given param is not valid") // 400 Bad Request // 501 Not Implemented``
)
