package handler

import (
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
)

type TaskHandler struct {
	taskService port.TaskService
}

func NewTaskHandler(taskService port.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) Route(router fiber.Router) {
	router.Get("/:id/topics", h.GetTopicsByTaskID)
	router.Post("/", h.CreateTask)
	router.Put("/:id", h.UpdateTask)
	router.Delete("/:id", h.DeleteTask)
}

func (h *TaskHandler) CreateTask(c fiber.Ctx) error {
	req := new(presenter.TaskCreateRequest)

	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	createdTask, err := h.taskService.CreateTask(c.Context(), req.GroupID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewTaskCreateResponse(createdTask))
}

func (h *TaskHandler) GetTopicsByTaskID(c fiber.Ctx) error {
	req := new(presenter.TaskGetTopicsByTaskIDRequest)

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	topics, err := h.taskService.GetTopicsByTaskID(c.Context(), req.ID, req.Offset, req.Limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTaskGetTopicsByIDResponse(topics))
}

func (h *TaskHandler) UpdateTask(c fiber.Ctx) error {
	req := new(presenter.TaskUpdateRequest)

	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	updatedTask, err := h.taskService.UpdateTask(c.Context(), req.ID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTaskUpdateResponse(updatedTask))
}

func (h *TaskHandler) DeleteTask(c fiber.Ctx) error {
	req := new(presenter.TaskDeleteRequest)

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	err := h.taskService.DeleteTask(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
