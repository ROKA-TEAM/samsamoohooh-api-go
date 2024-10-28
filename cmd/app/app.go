package main

import (
	"context"
	"log"
	"samsamoohooh-go-api/internal/application"
	"samsamoohooh-go-api/internal/infra"
	"samsamoohooh-go-api/internal/infra/storage/mysql"
	"samsamoohooh-go-api/internal/router"
	"time"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Supply(".env.toml"),
		fx.Provide(
			router.NewRouter,
		),

		infra.ConfigModule,
		infra.LoggerModule,
		infra.StorageModule,
		infra.AuthenticationModule,
		infra.MiddlewareModule,
		infra.ValidatorModule,

		application.RepositoryModule,
		application.ServiceModule,
		application.HandlerModule,

		fx.Invoke(
			func(r *router.Router) {},
			func(mysql *mysql.MySQL) {
				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				if err := mysql.AutoMigration(ctx); err != nil {
					log.Panicf("failed to migrate: %v", err)
				}
			},
		),
	).Run()
}
