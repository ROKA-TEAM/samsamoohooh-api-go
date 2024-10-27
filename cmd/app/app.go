package main

import (
	"samsamoohooh-go-api/internal/application"
	"samsamoohooh-go-api/internal/infra"
	"samsamoohooh-go-api/internal/router"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Supply(".env.toml"),
		fx.Provide(
			router.NewRouter,
		),

		infra.ConfigModule,
		infra.StorageModule,
		infra.AuthenticationModule,
		infra.MiddlewareModule,

		application.RepositoryModule,
		application.ServiceModule,
		application.HandlerModule,

		fx.Invoke(func(r *router.Router) {}),
	).Run()
}
