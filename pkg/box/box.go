package box

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
)

var errorsMap map[error]int = make(map[error]int)

func NewError(msg string, status int) error {
	newErr := errors.New(msg)
	errorsMap[newErr] = status
	return newErr
}

func GetStatus(err error) (int, bool) {
	causeErr := errors.Cause(err)
	got, ok := errorsMap[causeErr]
	return got, ok
}

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func Equal(err, target error) bool {
	return errors.Is(err, target)
}

func AppendMsg(err error, msg string) error {
	return errors.WithMessage(err, msg)
}

func GetFiberErrorHandler() func(c fiber.Ctx, err error) error {
	return func(c fiber.Ctx, err error) error {
		status, ok := GetStatus(err)
		if !ok {
			return fiber.DefaultErrorHandler(c, err)
		}
		return c.Status(status).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
}
