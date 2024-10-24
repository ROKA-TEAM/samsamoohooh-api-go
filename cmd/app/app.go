package main

import (
	"samsamoohooh-go-api/internal/application/handler"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/infra/authentication/token/jwt"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/storage/mysql"
	"samsamoohooh-go-api/internal/infra/storage/redis"
	"samsamoohooh-go-api/internal/router"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Supply(".env.toml"),
		fx.Provide(
			config.NewConfig,
			router.NewRouter,

			// database
			mysql.NewMySQL,

			// repositories
			fx.Annotate(
				redis.NewRedis,
				fx.As(new(port.RedisRepository)),
			),

			// services
			fx.Annotate(
				jwt.NewJWTService,
				fx.As(new(port.TokenService)),
			),

			// handlers
			handler.NewErrorHandler,
		),
		fx.Invoke(func(r *router.Router) {}),
	).Run()
}
