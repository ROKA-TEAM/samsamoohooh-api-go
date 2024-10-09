package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pkg/errors"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/port"
	"strings"
)

type GuardMiddleware struct {
	jwtService port.JWTService
}

func NewGuardMiddleware(jwtService port.JWTService) *GuardMiddleware {
	return &GuardMiddleware{jwtService: jwtService}
}

func (g *GuardMiddleware) getTokenStringInHeader(c fiber.Ctx) (string, error) {
	// header 에서 token을 가져온다.
	authHeader := c.Get("Authorization")
	if len(authHeader) == 0 {
		return "", errors.New("authorization header is empty")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("authorization header is not bearer token")
	}

	return strings.TrimPrefix(authHeader, "Bearer "), nil

}

func (g *GuardMiddleware) ProtectFromTempToken(c fiber.Ctx) error {

	tokenString, err := g.getTokenStringInHeader(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	_, err = g.jwtService.VerifyToken(tokenString)
	if errors.Is(err, domain.ErrTokenTemporary) {
		return c.Status(fiber.StatusUnauthorized).SendString("token is temporary")
	} else if err != nil {
		return err
	}

	return c.Next()
}

func (g *GuardMiddleware) AllowEntryOnlyTempToken(c fiber.Ctx) error {
	tokenString, err := g.getTokenStringInHeader(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	tokenPayload, err := g.jwtService.VerifyTempToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	fiber.Locals[domain.TempTokenPayload](c, domain.Temp, *tokenPayload)

	return c.Next()

}
