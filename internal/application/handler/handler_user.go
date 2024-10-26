package handler

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/presenter"
	"samsamoohooh-go-api/internal/infra/middleware/guard"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Route(router fiber.Router) {
	router.Get("/me", h.GetByMe)
	router.Get("/me/groups", h.GetGroupsByMe)
	router.Put("/me", h.UpdateUserByMe)
	router.Delete("/me", h.DeleteUserByMe)
}

func (h *UserHandler) GetByMe(c fiber.Ctx) error {
	token := fiber.Locals[*domain.Token](c, guard.TokenKey)

	gotUser, err := h.userService.GetByUserID(c.Context(), token.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserGetByMeResponse(gotUser))
}

func (h *UserHandler) GetGroupsByMe(c fiber.Ctx) error {
	token := fiber.Locals[*domain.Token](c, guard.TokenKey)

	req := new(presenter.UserGetGroupsByMe)
	if err := c.Bind().Query(req); err != nil {
		return err
	}

	gotGroups, err := h.userService.GetGroupsByUserID(c.Context(), token.ID, req.Limit, req.Offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserGetGroupsByMeResponse(gotGroups))
}

func (h *UserHandler) UpdateUserByMe(c fiber.Ctx) error {
	token := fiber.Locals[*domain.Token](c, guard.TokenKey)

	req := new(presenter.UserUpdateByMeRequest)
	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	udpatedUser, err := h.userService.UpdateUser(c.Context(), token.ID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserUpdateByMeResponse(udpatedUser))
}

func (h *UserHandler) DeleteUserByMe(c fiber.Ctx) error {
	token := fiber.Locals[*domain.Token](c, guard.TokenKey)

	err := h.userService.DeleteUser(c.Context(), token.ID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
