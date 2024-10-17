package token

import (
	"github.com/pkg/errors"
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
		return "", errors.Wrap(domain.ErrTokenGenerate, err.Error())
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
		return "", errors.Wrap(domain.ErrTokenGenerate, err.Error())
	}

	return tokenString, nil
}
func (s *JWTService) ValidateToken(tokenString string) (bool, error) {
	customClaims := new(customClaims)
	_, err := jwt.ParseWithClaims(tokenString, customClaims, func(t *jwt.Token) (any, error) {
		return []byte(s.config.Token.SecretKey), nil
	})
	if err != nil {
		return false, errors.Wrap(domain.ErrTokenParse, err.Error())
	}

	// 현재 시각
	now := time.Now()

	// 해석한 토큰의 issure가 일치하는가?
	if customClaims.Issuer != s.config.Token.Issuer {
		return false, domain.ErrInvalidTokenIssuer
	}

	// 해석한 토큰의 expiresAt이 유효한가? (현재 시간이 expiresAt보다 앞서 있다면)
	if now.After(customClaims.ExpiresAt.Time) {
		// TODO: warp error
		return false, domain.ErrTokenExpired
	}

	// 해석한 토큰의 notBefore가 유효한가? (현재 시간이 notBefore보다 않다면)
	if now.Before(customClaims.NotBefore.Time) {
		return false, domain.ErrTokenNotActiveYet
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
		return nil, errors.Wrap(domain.ErrTokenParse, err.Error())
	}

	return customClaims.toDomain(), nil
}
