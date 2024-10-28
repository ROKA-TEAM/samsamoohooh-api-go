package router

import (
	"context"
	"samsamoohooh-go-api/internal/application/handler"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/middleware/guard"
	"samsamoohooh-go-api/internal/infra/validator"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

type Router struct {
	config *config.Config
	app    *fiber.App

	// handlers
	oauthHandler   *handler.OauthHandler
	authHandler    *handler.AuthHandler
	userHandler    *handler.UserHandler
	groupHandler   *handler.GroupHandler
	postHandler    *handler.PostHandler
	commentHandler *handler.CommentHandler
	taskHandler    *handler.TaskHandler
	toicpHandler   *handler.TopicHandler

	// middleware
	guardMiddleware *guard.GuardMiddleware

	// validator
	validator *validator.Validator
}

func NewRouter(
	lc fx.Lifecycle,
	config *config.Config,

	// handler dependency
	errorHandler *handler.ErrorHandler,
	oauthHandler *handler.OauthHandler,
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	groupHandler *handler.GroupHandler,
	postHandler *handler.PostHandler,
	commentHandler *handler.CommentHandler,
	taskHandler *handler.TaskHandler,
	toicpHandler *handler.TopicHandler,

	// middleware
	guardMiddleware *guard.GuardMiddleware,

	// validator
	validator *validator.Validator,
) *Router {
	r := &Router{
		config: config,
		// handlers
		oauthHandler:   oauthHandler,
		authHandler:    authHandler,
		userHandler:    userHandler,
		groupHandler:   groupHandler,
		postHandler:    postHandler,
		commentHandler: commentHandler,
		taskHandler:    taskHandler,
		toicpHandler:   toicpHandler,

		// middleware
		guardMiddleware: guardMiddleware,

		// validator
		validator: validator,

		// init fiber app
		app: fiber.New(fiber.Config{
			ErrorHandler:    errorHandler.HandleError(),
			StructValidator: validator,
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
		api := v1.Group("/api")
		{
			oauth := api.Group("/oauth")
			{
				r.oauthHandler.Route(oauth)
			}

			auth := api.Group("/auth")
			{
				r.oauthHandler.Route(auth)
			}

			users := api.Group("/users", r.guardMiddleware.Authenticate)
			{
				r.userHandler.Route(users)
			}

			groups := api.Group("/groups", r.guardMiddleware.Authenticate)
			{
				r.groupHandler.Route(groups)
			}

			posts := api.Group("/posts", r.guardMiddleware.Authenticate)
			{
				r.postHandler.Route(posts)
			}

			comments := api.Group("/comments", r.guardMiddleware.Authenticate)
			{
				r.commentHandler.Route(comments)
			}

			tasks := api.Group("/tasks", r.guardMiddleware.Authenticate)
			{
				r.taskHandler.Route(tasks)
			}

			topics := api.Group("/topics", r.guardMiddleware.Authenticate)
			{
				r.toicpHandler.Route(topics)
			}
		}
	}
}

func (r *Router) SetMiddleware() {
}

func (r *Router) Start() error {
	return r.app.Listen(r.config.Server.Port)
}
