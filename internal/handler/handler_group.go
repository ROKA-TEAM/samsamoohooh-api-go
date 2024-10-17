package handler

import (
	"samsamoohooh-go-api/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type GroupHandler struct {
	groupService domain.GroupService
}

func NewGroupHandler(groupService domain.GroupService) *GroupHandler {
	return &GroupHandler{
		groupService: groupService,
	}
}

func (h *GroupHandler) Route(router fiber.Router) {

}

func (h *GroupHandler) Create(c *fiber.Ctx) error {
	return nil
}
