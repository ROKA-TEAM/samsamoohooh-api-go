package handler

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/presenter"
	"samsamoohooh-go-api/internal/infra/middleware/guard"

	"github.com/gofiber/fiber/v3"
)

type GroupHandler struct {
	groupService port.GroupService
}

func NewGroupHandler(groupService port.GroupService) *GroupHandler {
	return &GroupHandler{
		groupService: groupService,
	}
}

func (h *GroupHandler) Route(r fiber.Router) {
	r.Post("/", h.CreateGroup)
	r.Get("/:id/users", h.GetUsersByGroupID)
	r.Get("/:id/posts", h.GetPostsByGroupID)
	r.Get("/:id/tasks", h.GetTasksByGroupID)
	r.Put("/:id", h.UpdateGroup)
	r.Post("/:id/leave", h.GroupLeave)
	r.Post("/:id/join-code/generate", h.GroupGenerateJoinCode)
	r.Post("/join/:code", h.JoinGroupByCode)
	r.Post(":id/discussion/start", h.StartDiscussion)
}

func (h *GroupHandler) CreateGroup(c fiber.Ctx) error {
	token := fiber.Locals[*domain.Token](c, guard.TokenKey)

	req := new(presenter.GroupCreateReqeust)
	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	createdGroup, err := h.groupService.CreateGroup(c.Context(), token.ID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewGroupCreateResponse(createdGroup))
}

func (h *GroupHandler) GetUsersByGroupID(c fiber.Ctx) error {
	req := new(presenter.GroupGetUsersByGroupIDRequest)

	if err := c.Bind().Query(req); err != nil {
		return err
	}

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	gotUsers, err := h.groupService.GetUsersByGroupID(c.Context(), req.ID, req.Limit, req.Offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetUsersByGroupIDResponse(gotUsers))
}

func (h *GroupHandler) GetPostsByGroupID(c fiber.Ctx) error {
	req := new(presenter.GroupGetPostsByGroupIDRequest)

	if err := c.Bind().Query(req); err != nil {
		return err
	}

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	gotPosts, err := h.groupService.GetPostsByGroupID(c.Context(), req.ID, req.Limit, req.Offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetPostsByGroupIDResponse(gotPosts))
}

func (h *GroupHandler) GetTasksByGroupID(c fiber.Ctx) error {
	req := new(presenter.GroupGetTasksByGroupIDRequest)

	if err := c.Bind().Query(req); err != nil {
		return err
	}

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	gotTasks, err := h.groupService.GetTasksByGroupID(c.Context(), req.ID, req.Limit, req.Offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetTasksByIDResponse(gotTasks))
}

func (h *GroupHandler) UpdateGroup(c fiber.Ctx) error {

	req := new(presenter.GroupUpdateRequest)
	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	updatedGroup, err := h.groupService.UpdateGroup(c.Context(), req.ID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupUpdateResponse(updatedGroup))
}

func (h *GroupHandler) GroupLeave(c fiber.Ctx) error {
	token := fiber.Locals[*domain.Token](c, guard.TokenKey)

	req := new(presenter.GroupLeaveRequest)
	if err := c.Bind().URI(req); err != nil {
		return err
	}

	err := h.groupService.LeaveGroup(c.Context(), req.ID, token.ID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *GroupHandler) GroupGenerateJoinCode(c fiber.Ctx) error {
	req := new(presenter.GroupGenerateJoinCodeRequest)

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	joinCode, err := h.groupService.GenerateJoinCode(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGenerateJoinCodeResponse(joinCode))
}

func (h *GroupHandler) JoinGroupByCode(c fiber.Ctx) error {
	token := fiber.Locals[*domain.Token](c, guard.TokenKey)

	req := new(presenter.GroupJoinByCodeRequest)
	if err := c.Bind().URI(req); err != nil {
		return err
	}

	err := h.groupService.JoinGroupByCode(c.Context(), token.ID, req.Code)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *GroupHandler) StartDiscussion(c fiber.Ctx) error {
	req := new(presenter.GroupStartDiscussionRequest)

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	if err := c.Bind().Body(req); err != nil {
		return err
	}

	topics, names, err := h.groupService.StartDiscussion(c.Context(), req.ID, req.TaskID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupStartDiscussionResponse(topics, names))
}
