package jwt

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var _ port.TokenService = (*JWTService)(nil)

type JWTService struct {
	config *config.Config
}

func NewJWTService(config *config.Config) *JWTService {
	return &JWTService{
		config: config,
	}
}

func (s *JWTService) GenerateAccessTokenString(id int, role domain.UserRoleType) (string, error) {
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
		ID:   id,
		Role: role,
		Type: domain.TokenTypeAccess,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString([]byte(s.config.Token.SecretKey))
	if err != nil {
		return "", exception.New(
			err,
			exception.ErrJWTSignature,
			exception.StatusInternalServerError,
			"Failed to sign Access token",
		)
	}

	return tokenString, nil
}

func (s *JWTService) GenerateRefreshTokenString(id int, role domain.UserRoleType) (string, error) {
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
		ID:   id,
		Role: role,
		Type: domain.TokenTypeRefresh,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString([]byte(s.config.Token.SecretKey))
	if err != nil {
		return "", exception.New(
			err,
			exception.ErrJWTSignature,
			exception.StatusInternalServerError,
			"Failed to sign Refresh token",
		)
	}

	return tokenString, nil
}

func (s *JWTService) ParseToken(tokenString string) (*domain.Token, error) {
	customClaims := new(customClaims)
	_, err := jwt.ParseWithClaims(tokenString, customClaims, func(t *jwt.Token) (any, error) {
		return []byte(s.config.Token.SecretKey), nil
	})
	if err != nil {
		return nil, exception.New(
			err,
			exception.ErrJWTParseFailed,
			exception.StatusUnauthorized,
			"Failed to parse token",
		)
	}
	now := time.Now()

	// 해석한 토큰의 issure가 일치하는가?
	if customClaims.Issuer != s.config.Token.Issuer {
		return nil, exception.NewWithoutErr(
			exception.ErrJWTInvalidIssuer,
			exception.StatusUnauthorized,
			"Invalid issuer",
		)
	}

	// 해석한 토큰의 expiresAt이 유효한가? (현재 시간이 expiresAt보다 앞서 있다면)
	if now.After(customClaims.ExpiresAt.Time) {
		return nil, exception.NewWithoutErr(
			exception.ErrJWTTokenExpired,
			exception.StatusUnauthorized,
			"Token expired",
		)
	}

	// 해석한 토큰의 notBefore가 유효한가? (현재 시간이 notBefore보다 않다면)
	if now.Before(customClaims.NotBefore.Time) {
		return nil, exception.NewWithoutErr(
			exception.ErrJWTNotActiveYet,
			exception.StatusUnauthorized,
			"Token not active yet",
		)
	}

	return customClaims.ToDomain(), nil
}
