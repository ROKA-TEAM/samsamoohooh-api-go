package guard

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/infra/exception"
	"strings"

	"github.com/gofiber/fiber/v3"
)

const (
	TokenKey = "token"
)

type GuardMiddleware struct {
	tokenService port.TokenService
}

func NewGuardMiddleware(tokenService port.TokenService) *GuardMiddleware {
	return &GuardMiddleware{
		tokenService: tokenService,
	}
}

func (g *GuardMiddleware) Authenticate(c fiber.Ctx) error {
	authorizationHeader := c.Get("Authorization")
	if authorizationHeader == "" {
		return exception.NewWithoutErr(
			exception.ErrMiddlewareGuardAuthorizatioHeaderMissing,
			exception.StatusUnauthorized,
			"Authorization header is missing",
		)
	}

	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
	if tokenString == "" {
		return exception.NewWithoutErr(
			exception.ErrMiddlewareGuardAuthorizatioHeaderMissing,
			exception.StatusUnauthorized,
			"Bearer token is missing",
		)
	}

	parsedToken, err := g.tokenService.ParseToken(tokenString)
	if err != nil {
		return err
	}

	if parsedToken.Type == domain.TokenTypeRefresh {
		return exception.NewWithoutErr(
			exception.ErrJWTInvalid,
			exception.StatusUnauthorized,
			"Refresh token is not allowed",
		)
	}

	c.Locals(TokenKey, parsedToken)
	return c.Next()
}
