package jwt

import (
	"samsamoohooh-go-api/internal/application/domain"

	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	jwt.RegisteredClaims
	ID   int                 `json:"id"`
	Role domain.UserRoleType `json:"role"`
	Type domain.TokenType    `json:"type"`
}

func (c customClaims) ToDomain() *domain.Token {
	return &domain.Token{
		Issuer:    c.Issuer,
		ExpiresAt: c.ExpiresAt.Time,
		NotBefore: c.NotBefore.Time,
		IssuedAt:  c.IssuedAt.Time,
		ID:        c.ID,
		Role:      c.Role,
		Type:      c.Type,
	}
}
