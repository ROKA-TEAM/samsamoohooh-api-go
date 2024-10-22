package domain

import (
	"net/http"
	"samsamoohooh-go-api/pkg/box"
)

var (
	// database errors
	ErrConstraint    = box.NewError("constraint violation", http.StatusBadRequest)
	ErrNotFound      = box.NewError("not found", http.StatusNotFound)
	ErrNotLoaded     = box.NewError("not loaded", http.StatusNotFound)
	ErrNotSingular   = box.NewError("not singular", http.StatusNotFound)
	ErrValidation    = box.NewError("validation error", http.StatusBadRequest)
	ErrAuthorization = box.NewError("authorization error", http.StatusUnauthorized)
	ErrForbidden     = box.NewError("forbidden", http.StatusForbidden)
	ErrInternal      = box.NewError("internal error", http.StatusInternalServerError)
	ErrBadRequest    = box.NewError("bad request", http.StatusBadRequest)
)
