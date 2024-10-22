package handler

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/handler/utils"
	"samsamoohooh-go-api/internal/application/presenter"
	"samsamoohooh-go-api/internal/infra/middleware/guard"
	"samsamoohooh-go-api/pkg/token"

	"github.com/gofiber/fiber/v3"
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
	router.Post("/", h.CreateGroup)
	router.Get("/:gid", h.GetByGroupID)
	router.Get("/:gid/users", h.GetUsersByGroupID)
	router.Get("/:gid/posts", h.GetPostsByGroupID)
	router.Get("/:gid/tasks", h.GetTasksByGroupID)
	router.Put("/:gid", h.UpdateGroup)
	router.Post("/:gid/tasks/:tid/discussion/start", h.StartDiscussion)

	router.Post("/:gid/join-code/generate", h.GenerateJoinCode)
	router.Post("/join/:code", h.JoinGroup)
	router.Post("/:gid/leave", h.LeaveGroup)
}

func (h *GroupHandler) CreateGroup(c fiber.Ctx) error {
	token, err := utils.GetToken(c)
	if err != nil {
		return err
	}

	body := new(presenter.GroupCreateReqeust)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	createdGroup, err := h.groupService.CreateGroup(c.Context(), token.Subject, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewGroupCreateResponse(createdGroup))
}

func (h *GroupHandler) GetByGroupID(c fiber.Ctx) error {
	gid := fiber.Params[int](c, "gid")

	gotGroup, err := h.groupService.GetByGroupID(c.Context(), gid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetByIDResponse(gotGroup))
}

func (h *GroupHandler) GetUsersByGroupID(c fiber.Ctx) error {
	gid := fiber.Params[int](c, "gid")
	limit := fiber.Query[int](c, "limit", DefaultLimit)
	offset := fiber.Query[int](c, "offset", DefaultOffset)

	users, err := h.groupService.GetUsersByGroupID(c.Context(), gid, offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetUsersByIDResponse(users))
}

func (h *GroupHandler) GetPostsByGroupID(c fiber.Ctx) error {
	gid := fiber.Params[int](c, "gid")
	limit := fiber.Query[int](c, "limit", DefaultLimit)
	offset := fiber.Query[int](c, "offset", DefaultOffset)

	listPost, err := h.groupService.GetPostsByGroupID(c.Context(), gid, offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetPostsByIDResponse(listPost))
}

func (h *GroupHandler) GetTasksByGroupID(c fiber.Ctx) error {
	gid := fiber.Params[int](c, "gid")
	limit := fiber.Query[int](c, "limit", DefaultLimit)
	offset := fiber.Query[int](c, "offset", DefaultOffset)

	listTask, err := h.groupService.GetTasksByGroupID(c.Context(), gid, offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGetTasksByIDResponse(listTask))
}

func (h *GroupHandler) UpdateGroup(c fiber.Ctx) error {
	gid := fiber.Params[int](c, "gid")

	body := new(presenter.GroupUpdateRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	updatedGroup, err := h.groupService.UpdateGroup(c.Context(), gid, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupUpdateResponse(updatedGroup))
}

func (h *GroupHandler) StartDiscussion(c fiber.Ctx) error {
	gid := fiber.Params[int](c, "gid")
	tid := fiber.Params[int](c, "tid")

	topics, userNames, err := h.groupService.StartDiscussion(c.Context(), gid, tid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupStartDiscussionResponse(topics, userNames))
}

func (h *GroupHandler) GenerateJoinCode(c fiber.Ctx) error {
	gid := fiber.Params[int](c, "gid")

	joinCode, err := h.groupService.GenerateJoinCode(c.Context(), gid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewGroupGenerateJoinCodeResponse(joinCode))
}

func (h *GroupHandler) JoinGroup(c fiber.Ctx) error {
	token := fiber.Locals[*token.Token](c, guard.TokenKey)
	code := fiber.Params[string](c, "code")

	err := h.groupService.JoinGroupByCode(c.Context(), token.Subject, code)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(nil)
}

func (h *GroupHandler) LeaveGroup(c fiber.Ctx) error {
	gid := fiber.Params[int](c, "gid")
	token := fiber.Locals[*token.Token](c, guard.TokenKey)

	err := h.groupService.LeaveGroup(c.Context(), gid, token.Subject)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)

}
