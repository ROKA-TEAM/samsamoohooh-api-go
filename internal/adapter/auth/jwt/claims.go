package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"samsamoohooh-go-api/internal/core/domain"
)

type customClaims struct {
	jwt.RegisteredClaims
	Social string `json:"social"`
	Type   string `json:"type:"`
	Role   string `json:"role"`
}

func (c *customClaims) toDomain() *domain.TokenPayload {
	return &domain.TokenPayload{
		Issuer:    c.Issuer,
		Subject:   c.Subject,
		Audience:  c.Audience[0],
		ExpiresAt: c.ExpiresAt.Time,
		NotBefore: c.NotBefore.Time,
		IssuedAt:  c.IssuedAt.Time,
		Social:    domain.SocialType(c.Social),
		Role:      domain.RoleType(c.Role),
		Type:      domain.TokenType(c.Type),
	}
}

type tempTokenClaims struct {
	jwt.RegisteredClaims
	Social string `json:"social,omitempty"`
	Type   string `json:"type,omitempty"`
}

func (c *tempTokenClaims) toDomain() *domain.TempTokenPayload {
	return &domain.TempTokenPayload{
		Issuer:    c.Issuer,
		Subject:   c.Subject,
		Audience:  c.Audience[0],
		ExpiresAt: c.ExpiresAt.Time,
		NotBefore: c.NotBefore.Time,
		IssuedAt:  c.IssuedAt.Time,
		Social:    domain.SocialType(c.Social),
		Type:      domain.TokenType(c.Type),
	}
}
