package domain

import "time"

type TokenType string

const (
	Access  TokenType = "ACCESS"
	Refresh TokenType = "REFRESH"
	Temp    TokenType = "TEMP"
)

type TokenPayload struct {
	// See https://github.com/golang-jwt/jwt/blob/main/registered_claims.go
	// Token을 발급한 주체
	Issuer string

	// Token의 대상인 사용자의 ID를 나타냅니다.
	Subject string

	// 어떤 서비스나 어플리케이션에 사용되어야 하는지를 지정합니다
	Audience string

	// token이 언제 만료되는지 나타냅니다.
	ExpiresAt time.Time

	// Token이 어느 시점 이후에만 유효한지
	NotBefore time.Time

	// Token이 언제 발급되었는지 나타냅니다
	IssuedAt time.Time

	// 로그인에 사용된 oauth 발급 기관
	Social SocialType

	Role RoleType

	// 어떤 토큰 타입인지
	Type TokenType
}

type TempTokenPayload struct {
	// Token을 발급한 주체
	Issuer string

	// Token의 대상인 사용자의 ID를 나타냅니다.
	Subject string

	// 어떤 서비스나 어플리케이션에 사용되어야 하는지를 지정합니다
	Audience string

	// token이 언제 만료되는지 나타냅니다.
	ExpiresAt time.Time

	// Token이 어느 시점 이후에만 유효한지
	NotBefore time.Time

	// Token이 언제 발급되었는지 나타냅니다
	IssuedAt time.Time

	Social SocialType
	Type   TokenType
}
