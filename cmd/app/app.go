package app

import (
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
		),
		fx.Invoke(func(r *router.Router) {}),
	).Run()
}
