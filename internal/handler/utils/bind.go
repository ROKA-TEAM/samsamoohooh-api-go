package utils

import (
	"github.com/gofiber/fiber/v2"
	"samsamoohooh-go-api/internal/infra/verifier"
)

func ParseAndVerify(c *fiber.Ctx, out any) error {
	if err := c.BodyParser(out); err != nil {
		return err
	}

	if err := verifier.Get().Verify(out); err != nil {
		return err
	}

	return nil
}
