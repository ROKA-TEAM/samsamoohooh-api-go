package handler

import (
	"errors"
	"fmt"
	"samsamoohooh-go-api/internal/application/presenter"
	"samsamoohooh-go-api/internal/infra/exception"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type ErrorHandler struct {
	logger *zap.Logger
}

func NewErrorHandler(
	logger *zap.Logger,
) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (h ErrorHandler) HandleError() func(c fiber.Ctx, err error) error {
	return func(c fiber.Ctx, err error) error {
		var excep *exception.Exception
		if errors.As(err, &excep) {
			// logging (error)
			if excep.Status == exception.StatusInternalServerError {
				fmt.Print("print")
				h.logger.Error(
					"error occurred",
					zap.String("type", excep.Type),
					zap.Int("status", excep.Status),
					zap.String("message", excep.Message),
					zap.Any("data", excep.Data),
				)
			} else if 400 <= excep.Status && excep.Status < 500 {
				// logging (warn)
				fmt.Print("print")
				h.logger.Warn(
					"warn occurred",
					zap.String("type", excep.Type),
					zap.Int("status", excep.Status),
					zap.String("message", excep.Message),
					zap.Any("data", excep.Data),
				)
			}

			// internal server error
			if excep.Status == exception.StatusInternalServerError {
				excep.Err = errors.New("hides information because an internal server error occurred.")
			}

			// it's myqsl error
			if exception.IsMySQLError(excep.Type) {
				excep.Err = errors.New("mysql errors are hidden for information security reasons.")
			}

			errResp := &presenter.ErrorResponse{
				Type:    excep.Type,
				Status:  excep.Status,
				Message: excep.Message,
				Data:    excep.Data,
			}

			if excep.Err != nil {
				errResp.Detail = excep.Err.Error()
			}

			return c.Status(excep.Status).JSON(errResp)
		}

		// it's fiber error
		code := exception.StatusInternalServerError
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
		}

		return c.Status(code).JSON(&presenter.ErrorResponse{
			Type:    exception.ErrWebServerInternal,
			Message: fiberErr.Message,
			Status:  code,
		})
	}
}
