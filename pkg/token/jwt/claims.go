package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"samsamoohooh-go-api/pkg/token"
)

type customClaims struct {
	jwt.RegisteredClaims
	Subject int        `json:"subject"`
	Role    string     `json:"role"`
	Type    token.Type `json:"type"`
}

func (c customClaims) ToToken() *token.Token {
	return &token.Token{
		Issuer:    c.Issuer,
		ExpiresAt: c.ExpiresAt.Time,
		NotBefore: c.NotBefore.Time,
		IssuedAt:  c.IssuedAt.Time,
		Subject:   c.Subject,
		Role:      c.Role,
		Type:      c.Type,
	}
}
