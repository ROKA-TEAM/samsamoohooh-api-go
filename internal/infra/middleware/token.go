package middleware

import (
	"samsamoohooh-go-api/internal/domain"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TokenMiddleware struct {
	tokenService domain.TokenService
}

func NewTokenMiddleware(tokenService domain.TokenService) *TokenMiddleware {
	return &TokenMiddleware{tokenService: tokenService}
}

func (m TokenMiddleware) RequireAuthorization(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return domain.ErrMissingAuthorizationHeader
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

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
