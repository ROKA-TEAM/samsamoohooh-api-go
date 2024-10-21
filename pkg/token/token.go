package token

import (
	"errors"
	"time"
)

var (
	ErrTokenParse         = errors.New("token parse error")
	ErrInvalidTokenIssuer = errors.New("invalid token issuer")
	ErrTokenExpired       = errors.New("token expired")
	ErrTokenNotActiveYet  = errors.New("token is not active yet")
)

type Type string

const (
	Access  Type = "ACCESS"
	Refresh Type = "REFRESH"
)

type Token struct {
	Issuer    string    // 토큰을 발급한 주체(발행자), 일반적으로 발급하는 서버의 이름이나 URL
	ExpiresAt time.Time // 토큰의 만료 시간, 이 시간이 지나면 토큰은 더 이상 유효하지 않음
	NotBefore time.Time // 이 시간 이후에만 토큰이 유효하게 시작되는 시간 (토큰의 유효 시작 시간)
	IssuedAt  time.Time // 토큰이 발행된 시간, 토큰이 언제 생성되었는지를 나타냄

	Subject int    // 토큰의 주제(해당 토큰이 인증하는 사용자나 주체의 ID)
	Role    string // 사용자에게 할당된 역할 (예: 관리자, 일반 사용자)
	Type    Type   // 사용자에게 할당된 타입 (예: ACCESS, REFRESH)
}

type Service interface {
	GenerateAccessTokenString(subject int, role string) (string, error)
	GenerateRefreshTokenString(subject int, role string) (string, error)
	ValidateToken(tokenString string) (bool, error)
	ParseToken(tokenString string) (*Token, error)
}
