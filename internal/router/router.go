package router

import (
	"context"
	"samsamoohooh-go-api/internal/application/handler"
	"samsamoohooh-go-api/internal/infra/config"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

type Router struct {
	config *config.Config

	app *fiber.App
}

func NewRouter(
	lc fx.Lifecycle,
	config *config.Config,

	// handler dependency
	errorHandler *handler.ErrorHandler,
) *Router {
	r := &Router{
		config: config,

		// init fiber app
		app: fiber.New(fiber.Config{
			ErrorHandler: errorHandler.HandleError(),
		}),
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					r.Route()
					if err := r.Start(); err != nil {
						panic(err)
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return r.app.Shutdown()
			},
		},
	)

	return r
}

func (r *Router) Route() {

}

func (r *Router) Start() error {
	return r.app.Listen(r.config.Server.Port)
}
