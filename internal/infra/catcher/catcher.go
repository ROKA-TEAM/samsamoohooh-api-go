package catcher

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/infra/logger"
)

var ErrorHandler = func(c *fiber.Ctx, caughtErr error) error {
	var status int

	switch {
	case errors.Is(domain.ErrTokenGenerate, caughtErr):
		status = fiber.StatusInternalServerError

	case errors.Is(domain.ErrTokenParse, caughtErr):
		status = fiber.StatusUnauthorized

	case errors.Is(domain.ErrInvalidTokenIssuer, caughtErr):
		status = fiber.StatusUnauthorized

	case errors.Is(domain.ErrTokenNotActiveYet, caughtErr):
		status = fiber.StatusUnauthorized

	case errors.Is(domain.ErrTokenExpired, caughtErr):
		status = fiber.StatusUnauthorized

	case errors.Is(domain.ErrNotFound, caughtErr):
		status = fiber.StatusNotFound
	case errors.Is(domain.ErrNotLoaded, caughtErr):
		status = fiber.StatusBadRequest
	case errors.Is(domain.ErrConstraint, caughtErr):
		status = fiber.StatusConflict
	case errors.Is(domain.ErrNotSingular, caughtErr):
		status = fiber.StatusBadRequest

	case errors.Is(domain.ErrValidation, caughtErr):
		status = fiber.StatusBadRequest

	case errors.Is(domain.ErrMissingAuthorizationHeader, caughtErr):
		status = fiber.StatusUnauthorized

	case errors.Is(domain.ErrInternal, caughtErr):
		status = fiber.StatusInternalServerError

	case errors.Is(domain.ErrNotMatchState, caughtErr):
		status = fiber.StatusUnauthorized

	default:
		return fiber.DefaultErrorHandler(c, caughtErr)
	}

	if status == fiber.StatusInternalServerError {
		logger.Get().Error("unintentional errors", zap.Error(caughtErr))
	}

	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(status).SendString(caughtErr.Error())
}
