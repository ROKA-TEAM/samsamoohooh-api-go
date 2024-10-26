package router

import (
	"context"
	"samsamoohooh-go-api/internal/application/handler"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/middleware/guard"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

type Router struct {
	config *config.Config
	app    *fiber.App

	// handlers
	oauthHandler *handler.OauthHandler
	authHandler  *handler.AuthHandler
	userHandler  *handler.UserHandler

	// middleware
	guardMiddleware *guard.GuardMiddleware
}

func NewRouter(
	lc fx.Lifecycle,
	config *config.Config,

	// handler dependency
	errorHandler *handler.ErrorHandler,
	oauthHandler *handler.OauthHandler,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,

	// middleware
	guardMiddleware *guard.GuardMiddleware,
) *Router {
	r := &Router{
		config: config,
		// handlers
		oauthHandler: oauthHandler,
		authHandler:  authHandler,
		userHandler:  userHandler,

		// middleware
		guardMiddleware: guardMiddleware,

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

			auth := app.Group("/auth")
			{
				r.oauthHandler.Route(auth)
			}

			users := app.Group("/users", r.guardMiddleware.Authenticate)
			{
				r.userHandler.Route(users)
			}
		}
	}
}

func (r *Router) Start() error {
	return r.app.Listen(r.config.Server.Port)
}
