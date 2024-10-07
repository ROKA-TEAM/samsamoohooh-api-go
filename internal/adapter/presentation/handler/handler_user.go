package handler

import (
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/core/dto"
	"samsamoohooh-go-api/internal/core/port"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Create(c fiber.Ctx) error {
	body := new(dto.UserCreateRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	u, err := h.userService.Create(c.Context(), body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(u)
}

func (h *UserHandler) GetByID(c fiber.Ctx) error {
	id := fiber.Params[uint](c, "id")

	u, err := h.userService.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(u)
}

func (h *UserHandler) GetGroupsByID(c fiber.Ctx) error {
	id := fiber.Params[uint](c, "id")

	groups, err := h.userService.GetGroupsByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(groups)
}

func (h *UserHandler) GetAll(c fiber.Ctx) error {
	limit := fiber.Query[int](c, "limit")
	skip := fiber.Query[int](c, "skip")

	users, err := h.userService.GetAll(c.Context(), limit, skip)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) Update(c fiber.Ctx) error {
	id := fiber.Params[uint](c, "id")
	body := new(dto.UserUpdateRequest)

	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	u, err := h.userService.Update(c.Context(), id, body)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(u)
}

func (h *UserHandler) Delete(c fiber.Ctx) error {
	id := fiber.Params[uint](c, "id")

	u, err := h.userService.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(u)
}
