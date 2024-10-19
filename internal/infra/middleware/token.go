package middleware

import (
	"samsamoohooh-go-api/internal/domain"
	"strings"

	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type GuardMiddleware struct {
	tokenService domain.TokenService
}

func NewGuardMiddleware(tokenService domain.TokenService) *GuardMiddleware {
	return &GuardMiddleware{tokenService: tokenService}
}

func (m GuardMiddleware) RequireAuthorization(c *fiber.Ctx) error {
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

func (m GuardMiddleware) RequireAccess(accessibleRoles ...domain.UserRoleType) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token, ok := c.Locals("token").(*domain.Token)
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
			return errors.Wrap(domain.ErrForbidden, "you are in possession of an inaccessible coin.")
		}

		return c.Next()
	}
}
