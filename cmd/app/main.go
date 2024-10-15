package main

import (
	"fmt"
	"log"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/token"
)

func main() {
	cfg, err := config.NewConfig(".env.toml")
	if err != nil {
		log.Panicf("failed to load config: %v\n", err)
	}

	var tokenService domain.TokenService = token.NewJWTService(cfg)

	accessTokenString, err := tokenService.GenerateAccessTokenString(1, domain.TokenRoleGuest)
	if err != nil {
		log.Fatalf("failed to generate access token string: %v\n", err)
	}

	refreshTokenString, err := tokenService.GenerateRefreshTokenString(1, domain.TokenRoleGuest)
	if err != nil {
		log.Fatalf("failed to generate refresh token string: %v\n", err)
	}

	accessToken, err := tokenService.ParseToken(accessTokenString)
	if err != nil {
		log.Fatalf("failed to parse access token: %v\n", err)
	}

	accessTokenIsValid, err := tokenService.ValidateToken(accessTokenString)
	if err != nil {
		log.Fatalf("failed to validate access token: %v\n", err)
	}

	refreshToken, err := tokenService.ParseToken(refreshTokenString)
	if err != nil {
		log.Fatalf("failed to parse refresh token: %v\n", err)
	}

	refreshTokenIsValid, err := tokenService.ValidateToken(refreshTokenString)
	if err != nil {
		log.Fatalf("failed to validate refresh token: %v\n", err)
	}

	fmt.Printf("parsed access token: %+v isValid: %v\nparsed refresh token: %+v isValid: %v\n", accessToken, accessTokenIsValid, refreshToken, refreshTokenIsValid)
}
