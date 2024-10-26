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
	app    *fiber.App

	// handlers
	oauthHandler *handler.OauthHandler
}

func NewRouter(
	lc fx.Lifecycle,
	config *config.Config,

	// handler dependency
	errorHandler *handler.ErrorHandler,
	oauthHandler *handler.OauthHandler,
) *Router {
	r := &Router{
		config: config,
		// handlers
		oauthHandler: oauthHandler,

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
	v1 := r.app.Group("/v1")
	{
		app := v1.Group("/app")
		{
			oauth := app.Group("/oauth")
			{
				r.oauthHandler.Route(oauth)
			}
		}
	}
}

func (r *Router) Start() error {
	return r.app.Listen(r.config.Server.Port)
}
