package jwt

import (
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/pkg/token"
	"samsamoohooh-go-api/pkg/token/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var _ token.Service = (*Service)(nil)

type Service struct {
	config *config.Config
}

func NewService(config *config.Config) *Service {
	return &Service{config: config}
}

func (s *Service) GenerateAccessTokenString(subject int, role string) (string, error) {
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
		Type:    token.Access,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString([]byte(s.config.Token.SecretKey))
	if err != nil {
		return "", utils.Wrap(err)
	}

	return tokenString, nil
}

func (s *Service) GenerateRefreshTokenString(subject int, role string) (string, error) {
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
		Type:    token.Refresh,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString([]byte(s.config.Token.SecretKey))
	if err != nil {
		return "", utils.Wrap(err)
	}

	return tokenString, nil
}
func (s *Service) ValidateToken(tokenString string) (bool, error) {
	customClaims := new(customClaims)
	_, err := jwt.ParseWithClaims(tokenString, customClaims, func(t *jwt.Token) (any, error) {
		return []byte(s.config.Token.SecretKey), nil
	})
	if err != nil {
		return false, utils.Wrap(err)
	}
	now := time.Now()

	// 해석한 토큰의 issure가 일치하는가?
	if customClaims.Issuer != s.config.Token.Issuer {
		return false, utils.Wrap(token.ErrInvalidTokenIssuer)
	}

	// 해석한 토큰의 expiresAt이 유효한가? (현재 시간이 expiresAt보다 앞서 있다면)
	if now.After(customClaims.ExpiresAt.Time) {
		return false, utils.Wrap(token.ErrTokenExpired)
	}

	// 해석한 토큰의 notBefore가 유효한가? (현재 시간이 notBefore보다 않다면)
	if now.Before(customClaims.NotBefore.Time) {
		return false, utils.Wrap(token.ErrTokenNotActiveYet)
	}

	return true, nil
}

func (s *Service) ParseToken(tokenString string) (*token.Token, error) {
	customClaims := new(customClaims)
	_, err := jwt.ParseWithClaims(tokenString, customClaims, func(t *jwt.Token) (any, error) {
		return []byte(s.config.Token.SecretKey), nil
	})

	if err != nil {
		return nil, utils.Wrap(err)
	}

	return customClaims.ToToken(), nil
}
