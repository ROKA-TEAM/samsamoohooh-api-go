package token

import (
	"errors"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/infra/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var _ domain.TokenService = (*JWTService)(nil)

type JWTService struct {
	config *config.Config
}

func NewJWTService(config *config.Config) *JWTService {
	return &JWTService{
		config,
	}
}

func (s *JWTService) GenerateAccessTokenString(subject int, role domain.TokenRoleType) (string, error) {
	now := time.Now()
	expiresAt := now.Add(time.Duration(s.config.Token.Duration.Access.ValidityPeriod))
	notBefore := now.Add(time.Duration(s.config.Token.Duration.Access.ActivationDelay))

	claims := customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.config.Token.Issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(notBefore),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Subject: subject,
		Role:    role,
		Type:    domain.TokenTypeAccess,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.Token.SecretKey))
	if err != nil {
		// TODO: wrap error
		return "", err
	}

	return tokenString, nil
}

func (s *JWTService) GenerateRefreshTokenString(subject int, role domain.TokenRoleType) (string, error) {
	now := time.Now()

	expiresAt := now.Add(time.Duration(s.config.Token.Duration.Refresh.ValidityPeriod))
	notBefore := now.Add(time.Duration(s.config.Token.Duration.Refresh.ActivationDelay))

	claims := customClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.config.Token.Issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(notBefore),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Subject: subject,
		Role:    role,
		Type:    domain.TokenTypeRefresh,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.Token.SecretKey))
	if err != nil {
		// TODO: warp error
		return "", err
	}

	return tokenString, nil
}
func (s *JWTService) ValidateToken(tokenString string) (bool, error) {
	customClaims := new(customClaims)
	_, err := jwt.ParseWithClaims(tokenString, customClaims, func(t *jwt.Token) (any, error) {
		return []byte(s.config.Token.SecretKey), nil
	})
	if err != nil {
		return false, err
	}

	// 현재 시각
	now := time.Now()

	// 해석한 토큰의 issure가 일치하는가?
	if customClaims.Issuer != s.config.Token.Issuer {
		// TODO: warp error
		return false, errors.New("invalid token issuer")
	}

	// 해석한 토큰의 expiresAt이 유효한가? (현재 시간이 expiresAt보다 앞서 있다면)
	if now.After(customClaims.ExpiresAt.Time) {
		// TODO: warp error
		return false, errors.New("token is expired")
	}

	// 해석한 토큰의 notBefore가 유효한가? (현재 시간이 notBefore보다 않다면)
	if now.Before(customClaims.NotBefore.Time) {
		// TODO: warp error
		return false, errors.New("token is not active yet")
	}

	// 해석한 토큰의 subject가 유효한가? (이건 일단 보류)

	return true, nil
}
func (s *JWTService) ParseToken(tokenString string) (*domain.Token, error) {
	customClaims := new(customClaims)
	_, err := jwt.ParseWithClaims(tokenString, customClaims, func(t *jwt.Token) (any, error) {
		return []byte(s.config.Token.SecretKey), nil
	})

	if err != nil {
		// TODO: warp error
		return nil, err
	}

	return customClaims.toDomain(), nil
}
