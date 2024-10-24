package main

import (
	"samsamoohooh-go-api/internal/application/handler"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/router"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Supply(".env.toml"),
		fx.Provide(
			config.NewConfig,
			router.NewRouter,

			// handlers
			handler.NewErrorHandler,
		),
		fx.Invoke(func(r *router.Router) {}),
	).Run()
}
