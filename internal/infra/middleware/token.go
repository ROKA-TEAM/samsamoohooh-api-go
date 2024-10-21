package middleware

import (
	domain2 "samsamoohooh-go-api/internal/application/domain"
	"strings"

	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type GuardMiddleware struct {
	tokenService domain2.TokenService
}

func NewGuardMiddleware(tokenService domain2.TokenService) *GuardMiddleware {
	return &GuardMiddleware{tokenService: tokenService}
}

func (m GuardMiddleware) RequireAuthorization(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return domain2.ErrMissingAuthorizationHeader
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

func (m GuardMiddleware) RequireAccess(accessibleRoles ...domain2.UserRoleType) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token, ok := c.Locals("token").(*domain2.Token)
		if !ok {
			return fiber.ErrUnauthorized
		}

		// only admin
		isAccessible := false
		for _, role := range accessibleRoles {
			if token.Role == role {
				isAccessible = true
				break
			}
		}

		if !isAccessible {
			return errors.Wrap(domain2.ErrForbidden, "you are in possession of an inaccessible coin.")
		}

		return c.Next()
	}
}
