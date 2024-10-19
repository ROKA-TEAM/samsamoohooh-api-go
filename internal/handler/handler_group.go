package handler

import (
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/handler/utils"
	"samsamoohooh-go-api/internal/infra/middleware"
	"samsamoohooh-go-api/internal/infra/presenter"

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

func (h *GroupHandler) Route(router fiber.Router, guard *middleware.GuardMiddleware) {
	router.Post("/", guard.RequireAccess(domain.UserRoleAdmin, domain.UserRoleGuest), h.Create)
	router.Get("/", guard.RequireAccess(domain.UserRoleAdmin), h.List)
	router.Get("/:id", guard.RequireAccess(domain.UserRoleAdmin, domain.UserRoleGuest), h.GetByID)
	router.Get("/:id/users", guard.RequireAccess(domain.UserRoleAdmin, domain.UserRoleGuest), h.GetUsersByID)
	router.Get("/:id/posts", guard.RequireAccess(domain.UserRoleAdmin, domain.UserRoleGuest), h.GetPostsByID)
	router.Get("/:id/tasks", guard.RequireAccess(domain.UserRoleAdmin, domain.UserRoleGuest), h.GetTasksByID)
	router.Put("/:id", guard.RequireAccess(domain.UserRoleAdmin, domain.UserRoleGuest), h.Update)
	router.Delete("/:id", guard.RequireAccess(domain.UserRoleAdmin), h.Delete)
}

func (h *GroupHandler) Create(c *fiber.Ctx) error {
	body := new(presenter.GroupCreateReqeust)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	createdGroup, err := h.groupService.Create(c.Context(), body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewGroupCreateResponse(createdGroup))
}

func (h *GroupHandler) List(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", DefaultLimit)
	offset := c.QueryInt("offset", DefaultOffset)

	groups, err := h.groupService.List(c.Context(), limit, offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupListResponse(groups))
}

func (h *GroupHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	gotGroup, err := h.groupService.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetByIDResponse(gotGroup))
}

func (h *GroupHandler) GetUsersByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	limit := c.QueryInt("limit", DefaultLimit)
	offset := c.QueryInt("offset", DefaultOffset)

	users, err := h.groupService.GetUsersByID(c.Context(), id, offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetUsersByIDResponse(users))
}

func (h *GroupHandler) GetPostsByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	limit := c.QueryInt("limit", DefaultLimit)
	offset := c.QueryInt("offset", DefaultOffset)

	listPost, err := h.groupService.GetPostsByID(c.Context(), id, offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetPostsByIDResponse(listPost))
}

func (h *GroupHandler) GetTasksByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	limit := c.QueryInt("limit", DefaultLimit)
	offset := c.QueryInt("offset", DefaultOffset)

	listTask, err := h.groupService.GetTasksByID(c.Context(), id, offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetTasksByIDResponse(listTask))
}

func (h *GroupHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	body := new(presenter.GroupUpdateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	updatedGroup, err := h.groupService.Update(c.Context(), id, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupUpdateResponse(updatedGroup))
}

func (h *GroupHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = h.groupService.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
