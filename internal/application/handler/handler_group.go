package handler

import (
	"github.com/gofiber/fiber/v2"
	domain2 "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/handler/utils"
	"samsamoohooh-go-api/internal/application/presenter/v1"
	"samsamoohooh-go-api/internal/infra/middleware"
)

type GroupHandler struct {
	groupService domain2.GroupService
}

func NewGroupHandler(groupService domain2.GroupService) *GroupHandler {
	return &GroupHandler{
		groupService: groupService,
	}
}

func (h *GroupHandler) Route(router fiber.Router, guard *middleware.GuardMiddleware) {
	router.Post("/", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.Create)
	router.Get("/", guard.RequireAccess(domain2.UserRoleAdmin), h.List)
	router.Get("/:id", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.GetByID)
	router.Get("/:id/users", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.GetUsersByID)
	router.Get("/:id/posts", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.GetPostsByID)
	router.Get("/:id/tasks", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.GetTasksByID)
	router.Put("/:id", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.Update)
	router.Delete("/:id", guard.RequireAccess(domain2.UserRoleAdmin), h.Delete)

	router.Get("/:gid/tasks/:tid/discussion/start", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.StartDiscussion)
}

func (h *GroupHandler) Create(c *fiber.Ctx) error {
	body := new(v1.GroupCreateReqeust)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	createdGroup, err := h.groupService.Create(c.Context(), body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(v1.NewGroupCreateResponse(createdGroup))
}

func (h *GroupHandler) List(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", DefaultLimit)
	offset := c.QueryInt("offset", DefaultOffset)

	groups, err := h.groupService.List(c.Context(), limit, offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(v1.NewGroupListResponse(groups))
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

	return c.Status(fiber.StatusOK).JSON(v1.NewGroupGetByIDResponse(gotGroup))
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

	return c.Status(fiber.StatusOK).JSON(v1.NewGroupGetUsersByIDResponse(users))
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

	return c.Status(fiber.StatusOK).JSON(v1.NewGroupGetPostsByIDResponse(listPost))
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

	return c.Status(fiber.StatusOK).JSON(v1.NewGroupGetTasksByIDResponse(listTask))
}

func (h *GroupHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	body := new(v1.GroupUpdateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	updatedGroup, err := h.groupService.Update(c.Context(), id, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(v1.NewGroupUpdateResponse(updatedGroup))
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

func (h *GroupHandler) StartDiscussion(c *fiber.Ctx) error {
	gid, err := c.ParamsInt("gid")
	if err != nil {
		return err
	}

	tid, err := c.ParamsInt("tid")
	if err != nil {
		return err
	}

	topics, userNames, err := h.groupService.StartDiscussion(c.Context(), gid, tid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(v1.NewGruopStartDiscussionResponse(topics, userNames))
}
