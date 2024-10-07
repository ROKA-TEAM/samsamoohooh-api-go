package router

import (
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/adapter/presentation/router/validator"
	"samsamoohooh-go-api/internal/infra/config"
)

const DefaultStartPort = ":8080"

type Router struct {
	*fiber.App
	config     *config.Config
	handlerSet HandlerSet
}

func New(config *config.Config, handlerSet HandlerSet) *Router {

	r := &Router{config: config, handlerSet: handlerSet, App: fiber.New(fiber.Config{
		AppName:         config.HTTP.Name,
		StructValidator: validator.New(),
		ErrorHandler:    customErrorHandler,
	})}

	r.setMiddleware()
	r.route()

	return r
}

// set middleware
func (r *Router) setMiddleware() {

}

// Route
func (r *Router) route() {
	// users
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.Post("/", r.handlerSet.UserHandler.Create)
				users.Get("/:id", r.handlerSet.UserHandler.GetByID)
				users.Get("/:id/groups", r.handlerSet.UserHandler.GetGroupsByID)
				users.Get("/", r.handlerSet.UserHandler.GetAll)
				users.Put("/:id", r.handlerSet.UserHandler.Update)
				users.Delete("/:id", r.handlerSet.UserHandler.Delete)
			}
		}
	}
}

func (r *Router) Start() error {
	startPort := DefaultStartPort
	if r.config.HTTP.Port != "" {
		startPort = r.config.HTTP.Port
	}

	return r.Listen(startPort)
}
