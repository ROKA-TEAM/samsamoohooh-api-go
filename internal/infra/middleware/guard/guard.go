package guard

import (
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/pkg/token"
	"strings"
)

const (
	TokenKey = "guard_token"
)

type Middleware struct {
	tokenService token.Service
	userService  domain.UserService
}

func NewMiddleware(
	tokenService token.Service,
	userService domain.UserService,
) *Middleware {
	return &Middleware{
		tokenService: tokenService,
		userService:  userService,
	}
}

func (m *Middleware) RequireAuthorization(c fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
	}

	tokenString = strings.TrimSpace(tokenString)
	isValid, err := m.tokenService.ValidateToken(tokenString)
	if err != nil {
		return err
	}

	if !isValid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid Authorization header")
	}

	parsedToken, err := m.tokenService.ParseToken(tokenString)
	if err != nil {
		return err
	}

	c.Locals(TokenKey, parsedToken)
	return c.Next()
}

func (m *Middleware) AccessOnly(roleType domain.UserRoleType) func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		t, ok := c.Locals(TokenKey).(*token.Token)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "unable to get token")
		}

		if !(strings.Compare(t.Role, string(roleType)) == 0) {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid role")
		}

		return c.Next()
	}
}

//func (m GuardMiddleware) RequireAccess(accessibleRoles ...domain2.UserRoleType) func(*fiber.Ctx) error {
//	return func(c *fiber.Ctx) error {
//		token, ok := c.Locals("token").(*domain2.Token)
//		if !ok {
//			return fiber.ErrUnauthorized
//		}
//
//		// only admin
//		isAccessible := false
//		for _, role := range accessibleRoles {
//			if token.Role == role {
//				isAccessible = true
//				break
//			}
//		}
//
//		if !isAccessible {
//			return errors.Wrap(domain2.ErrForbidden, "you are in possession of an inaccessible coin.")
//		}
//
//		return c.Next()
//	}
//}
