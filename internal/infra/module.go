package infra

import (
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/infra/authentication/oauth/google"
	"samsamoohooh-go-api/internal/infra/authentication/oauth/kakao"
	"samsamoohooh-go-api/internal/infra/authentication/token/jwt"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/logger/zap"
	"samsamoohooh-go-api/internal/infra/middleware/guard"
	"samsamoohooh-go-api/internal/infra/storage/mysql"
	"samsamoohooh-go-api/internal/infra/storage/redis"
	"samsamoohooh-go-api/internal/infra/validator"

	"go.uber.org/fx"
)

var StorageModule = fx.Module(
	"storage-module",
	fx.Provide(
		mysql.NewMySQL,
		fx.Annotate(
			redis.NewRedis,
			fx.As(new(port.RedisRepository)),
		),
	),
)

var MiddlewareModule = fx.Module(
	"middleware-module",
	fx.Provide(
		guard.NewGuardMiddleware,
	),
)

var ConfigModule = fx.Module(
	"config-module",
	fx.Provide(
		config.NewConfig,
	),
)

var AuthenticationModule = fx.Module(
	"authentication-module",
	fx.Provide(
		fx.Annotate(
			jwt.NewJWTService,
			fx.As(new(port.TokenService)),
		),

		fx.Annotate(
			google.NewGoogleOauthService,
			fx.As(new(port.OauthImplictGrantService)),
			fx.ResultTags(`name:"google"`),
		),
		fx.Annotate(
			kakao.NewKakaoOauthService,
			fx.As(new(port.OauthImplictGrantService)),
			fx.ResultTags(`name:"kakao"`),
		),
	),
)

var ValidatorModule = fx.Module(
	"validator-module",
	fx.Provide(
		validator.NewValidator,
	),
)

var LoggerModule = fx.Module(
	"logger-module",
	fx.Provide(
		zap.NewCustomZapLogger,
	),
)
