package token

import (
	"samsamoohooh-go-api/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type customClaims struct {
	jwt.RegisteredClaims
	Subject int                 `json:"subject"`
	Role    domain.UserRoleType `json:"role"`
	Type    domain.TokenType    `json:"type"`
}

func (c customClaims) toDomain() *domain.Token {
	return &domain.Token{
		Issuer:    c.Issuer,
		Subject:   c.Subject,
		ExpiresAt: c.ExpiresAt.Time,
		NotBefore: c.NotBefore.Time,
		IssuedAt:  c.IssuedAt.Time,
		Role:      c.Role,
		Type:      c.Type,
	}
}
