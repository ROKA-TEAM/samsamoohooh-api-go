package handler

import (
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/handler/utils"
	"samsamoohooh-go-api/internal/application/presenter"
	"samsamoohooh-go-api/internal/infra/middleware/guard"
)

type UserHandler struct {
	userService domain.UserService
}

func NewUserHandler(
	userService domain.UserService,
) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Route(router fiber.Router, guard *guard.Middleware) {
	me := router.Group("/me", guard.RequireAuthorization, guard.AccessOnly(domain.UserRoleUser))
	{
		me.Get("/", h.GetByMe)
		me.Get("/groups", h.GetGroupsByMe)
		me.Put("/", h.UpdateMe)
		me.Delete("/", h.DeleteMe)
	}
}

func (h *UserHandler) GetByMe(c fiber.Ctx) error {
	token, err := utils.GetToken(c)
	if err != nil {
		return err
	}

	gotUser, err := h.userService.GetByUserID(c.Context(), token.Subject)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserGetByMeResponse(gotUser))
}

func (h *UserHandler) GetGroupsByMe(c fiber.Ctx) error {
	token, err := utils.GetToken(c)
	if err != nil {
		return err
	}
	limit := fiber.Query[int](c, "limit")
	offset := fiber.Query[int](c, "offset")

	listGroup, err := h.userService.GetGroupsByUserID(c.Context(), token.Subject, limit, offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserGetGroupsByMeResponse(listGroup))
}

func (h *UserHandler) UpdateMe(c fiber.Ctx) error {
	token, err := utils.GetToken(c)
	if err != nil {
		return err
	}
	body := new(presenter.UserUpdateMeRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	updatedUser, err := h.userService.UpdateUser(c.Context(), token.Subject, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserUpdateMeResponse(updatedUser))
}

func (h *UserHandler) DeleteMe(c fiber.Ctx) error {
	token, err := utils.GetToken(c)
	if err != nil {
		return err
	}

	err = h.userService.DeleteUser(c.Context(), token.Subject)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
