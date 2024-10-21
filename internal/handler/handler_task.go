package handler

import (
	domain2 "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/handler/utils"
	"samsamoohooh-go-api/internal/infra/middleware"
	"samsamoohooh-go-api/internal/infra/presenter"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskService domain2.TaskService
}

func NewTaskHandler(taskService domain2.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) Route(router fiber.Router, guard *middleware.GuardMiddleware) {
	router.Post("/", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.Create)
	router.Get("/", guard.RequireAccess(domain2.UserRoleAdmin), h.List)
	router.Get("/:id", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.GetByID)
	router.Get("/:id/topics", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.GetTopicsByID)
	router.Put("/:id", guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest), h.Update)
	router.Delete("/", guard.RequireAccess(domain2.UserRoleAdmin), h.Delete)
}

func (h *TaskHandler) Create(c *fiber.Ctx) error {
	body := new(presenter.TaskCreateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	createdTask, err := h.taskService.Create(c.Context(), body.GroupID, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewTaskCreateResponse(createdTask))
}

func (h *TaskHandler) List(c *fiber.Ctx) error {
	offset := c.QueryInt("offset", DefaultOffset)
	limit := c.QueryInt("limit", DefaultLimit)

	listTask, err := h.taskService.List(c.Context(), offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTaskListResponse(listTask))
}

func (h *TaskHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	gotTask, err := h.taskService.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTaskGetByIDResponse(gotTask))
}

func (h *TaskHandler) GetTopicsByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	offset := c.QueryInt("offset", DefaultOffset)
	limit := c.QueryInt("limit", DefaultLimit)

	listTopics, err := h.taskService.GetTopicsByID(c.Context(), id, offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTaskGetTopicsByIDResponse(listTopics))
}

func (h *TaskHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	body := new(presenter.TaskUpdateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	updatedTask, err := h.taskService.Updated(c.Context(), id, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTaskUpdateResponse(updatedTask))
}

func (h *TaskHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = h.taskService.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
