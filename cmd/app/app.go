package main

import (
	"samsamoohooh-go-api/internal/application/handler"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/repository"
	"samsamoohooh-go-api/internal/application/service"
	"samsamoohooh-go-api/internal/infra/authentication/oauth/google"
	"samsamoohooh-go-api/internal/infra/authentication/oauth/kakao"
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

			fx.Annotate(
				repository.NewUserRepository,
				fx.As(new(port.UserRepository)),
			),

			// services
			fx.Annotate(
				jwt.NewJWTService,
				fx.As(new(port.TokenService)),
			),

			fx.Annotate(
				service.NewUserService,
				fx.As(new(port.UserService)),
			),

			fx.Annotate(
				google.NewGoogleOauthService,
				fx.As(new(port.ImplictGrantService)),
				fx.ResultTags(`name:"google"`),
			),

			fx.Annotate(
				kakao.NewKakaoOauthService,
				fx.As(new(port.ImplictGrantService)),
				fx.ResultTags(`name:"kakao"`),
			),
			handler.NewErrorHandler,
		),
		fx.Invoke(func(r *router.Router) {}),
	).Run()
}
