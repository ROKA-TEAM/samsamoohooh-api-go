package main

import (
	"log"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/pkg/token/jwt"
)

func main() {
	cfg, err := config.NewConfig(".env.toml")
	if err != nil {
		log.Panicf("failed to load config: %v\n", err)
	}

	jwtService := jwt.NewService(cfg)
	_ = jwtService

}
