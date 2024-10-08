package port

import "samsamoohooh-go-api/internal/core/domain"

type JWTService interface {
	CreateAccessToken(user *domain.User) (string, error)
	CreateRefreshToken(user *domain.User) (string, error)
	CreateTempToken(userID uint, sub, social string) (string, error)
	VerifyToken(tokenString string) (*domain.TokenPayload, error)
}
