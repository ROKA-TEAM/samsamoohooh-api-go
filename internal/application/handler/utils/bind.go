package utils

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/infra/verifier"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func ParseAndVerify(c *fiber.Ctx, out any) error {
	if err := c.BodyParser(out); err != nil {
		return errors.Wrap(domain.ErrParsing, err.Error())
	}

	if err := verifier.Get().Verify(out); err != nil {
		err := errors.Wrap(domain.ErrValidation, err.Error())
		return err
	}
	// 왜 에러가 warp 안되냐?
	return nil
}
