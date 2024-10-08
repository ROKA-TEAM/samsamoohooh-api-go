package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"samsamoohooh-go-api/internal/core/domain"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) GenerateSecureToken(length int) (string, error) {
	if length < 1 {
		return "", errors.Wrap(domain.ErrInternal, fmt.Sprintf("invalid token length: %d", length))
	}

	bytes := make([]byte, length)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes), nil
}
