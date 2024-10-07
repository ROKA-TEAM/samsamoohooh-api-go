package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"samsamoohooh-go-api/internal/core/domain"
)

var customErrorHandler = func(c fiber.Ctx, err error) error {

	var status int
	switch {
	case isValidationError(err):
		status = http.StatusBadRequest
	// domain error
	case errors.Is(err, domain.ErrUnauthorized):
		status = http.StatusUnauthorized
	case errors.Is(err, domain.ErrInternal):
		status = http.StatusInternalServerError
	// gorm error
	case errors.Is(err, gorm.ErrRecordNotFound):
		status = http.StatusNotFound
	case errors.Is(err, gorm.ErrDuplicatedKey):
		status = http.StatusConflict
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		status = http.StatusBadRequest
	case errors.Is(err, gorm.ErrCheckConstraintViolated):
		status = http.StatusBadRequest
	case errors.Is(err, gorm.ErrInvalidTransaction):
		status = http.StatusBadRequest
	case errors.Is(err, gorm.ErrInvalidData),
		errors.Is(err, gorm.ErrInvalidField),
		errors.Is(err, gorm.ErrInvalidValue),
		errors.Is(err, gorm.ErrInvalidValueOfLength):
		status = http.StatusBadRequest
	case errors.Is(err, gorm.ErrMissingWhereClause),
		errors.Is(err, gorm.ErrPrimaryKeyRequired),
		errors.Is(err, gorm.ErrModelValueRequired),
		errors.Is(err, gorm.ErrModelAccessibleFieldsRequired),
		errors.Is(err, gorm.ErrSubQueryRequired):
		status = http.StatusBadRequest
	case errors.Is(err, gorm.ErrNotImplemented),
		errors.Is(err, gorm.ErrUnsupportedRelation),
		errors.Is(err, gorm.ErrUnsupportedDriver),
		errors.Is(err, gorm.ErrDryRunModeUnsupported):
		status = http.StatusNotImplemented
	case errors.Is(err, gorm.ErrRegistered),
		errors.Is(err, gorm.ErrEmptySlice),
		errors.Is(err, gorm.ErrInvalidDB),
		errors.Is(err, gorm.ErrPreloadNotAllowed):
		status = http.StatusInternalServerError
	default:
		return fiber.DefaultErrorHandler(c, err)
	}

	// Set Content-Type: text/plain; charset=utf-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	// Return status code with error message
	return c.Status(status).SendString(err.Error())
}

func isValidationError(err error) bool {
	var validationErrors validator.ValidationErrors
	return errors.As(err, &validationErrors)
}
