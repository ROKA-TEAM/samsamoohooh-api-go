package port

import "samsamoohooh-go-api/internal/application/domain"

type TokenService interface {
	GenerateAccessTokenString(id int, role domain.UserRoleType) (string, error)
	GenerateRefreshTokenString(id int, role domain.UserRoleType) (string, error)
	ParseToken(tokenString string) (*domain.Token, error)
}
