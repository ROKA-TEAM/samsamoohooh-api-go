package router

import (
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/config"
)

type Router struct {
	app    *fiber.App
	config config.Config
	UserRouter
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) route() {
	api := r.app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			r.UserRouter.Route(v1)
		}
	}

}

func (r *Router) Start() error {
	return r.app.Listen(r.config.HTTP.Port)
}
