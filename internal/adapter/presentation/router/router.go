package router

import (
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/infra/config"
)

const DefaultStartPort = ":8080"

type Router struct {
	engine *fiber.App
	config *config.Config
}

func New(config *config.Config) *Router {

	r := &Router{config: config, engine: fiber.New(fiber.Config{})}

	r.setMiddleware()
	r.setRoute()

	return r
}

func (r *Router) Start() error {
	startPort := DefaultStartPort
	if r.config.HTTP.Port != "" {
		startPort = r.config.HTTP.Port
	}

	return r.engine.Listen(startPort)
}

// set middleware
func (r *Router) setMiddleware() {

}

// Route
func (r *Router) setRoute() {

}
