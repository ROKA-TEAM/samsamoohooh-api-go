package router

import (
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/adapter/presentation/http/handler"
)

type UserRouter struct {
	userHandler handler.UserHandler
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (UserRouter) Route(r fiber.Router) {
	users := r.Group("users")
	users.Get("/", func(ctx fiber.Ctx) error {
		return nil
	})
}
