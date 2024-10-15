package middleware

import (
	"samsamoohooh-go-api/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type TokenMiddleware struct {
	tokenService domain.TokenService
}

func (m TokenMiddleware) Authorization(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")[len("Bearer "):]

	isValid, err := m.tokenService.ValidateToken(tokenString)
	if !isValid || err != nil {
		return err
	}

	token, err := m.tokenService.ParseToken(tokenString)
	if err != nil {
		return err
	}

	c.Locals("token", token)
	return c.Next()
}
