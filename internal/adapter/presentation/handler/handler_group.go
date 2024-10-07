package handler

import (
	"github.com/gofiber/fiber/v3"
	"samsamoohooh-go-api/internal/core/dto"
	"samsamoohooh-go-api/internal/core/port"
)

type GroupHandler struct {
	groupService port.GroupService
}

func NewGroupHandler(groupService port.GroupService) *GroupHandler {
	return &GroupHandler{groupService: groupService}
}

func (g *GroupHandler) Create(c fiber.Ctx) error {
	body := new(dto.GroupCreateRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	createdGroup, err := g.groupService.Create(c.Context(), body, 6)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(createdGroup)
}

func (g *GroupHandler) GetByID(c fiber.Ctx) error {
	id := fiber.Params[uint](c, "id")

	queriedGroup, err := g.groupService.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(queriedGroup)
}

func (g *GroupHandler) GetUsersByID(c fiber.Ctx) error {
	id := fiber.Params[uint](c, "id")

	queriedUsers, err := g.groupService.GetUsersByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(queriedUsers)
}
