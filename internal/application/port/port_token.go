package port

import "samsamoohooh-go-api/internal/application/domain"

type TokenService interface {
	GenerateAccessTokenString(subject int, role string) (string, error)
	GenerateRefreshTokenString(subject int, role string) (string, error)
	ValidateToken(tokenString string) (bool, error)
	ParseToken(tokenString string) (*domain.Token, error)
}
