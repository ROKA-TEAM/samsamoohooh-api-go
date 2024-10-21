package handler

import (
	domain "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
)

type TaskHandler struct {
	taskService domain.TaskService
}

func NewTaskHandler(
	taskService domain.TaskService,
) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(c fiber.Ctx) error {
	body := new(presenter.TaskCreateRequest)

	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	createdTask, err := h.taskService.CreateTask(c.Context(), body.GroupID, body.ToDomain())
	if err != nil {
		return err
	}

	return c.JSON(presenter.NewTaskCreateResponse(createdTask))
}

func (h *TaskHandler) GetByTaskID(c fiber.Ctx) error {
	tid := fiber.Params[int](c, "tid")

	task, err := h.taskService.GetByTaskID(c.Context(), tid)
	if err != nil {
		return err
	}

	return c.JSON(presenter.NewTaskGetByTaskIDResponse(task))
}

func (h *TaskHandler) UpdateTask(c fiber.Ctx) error {
	tid := fiber.Params[int](c, "tid")
	body := new(presenter.TaskUpdateRequest)

	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	updatedTask, err := h.taskService.UpdateTask(c.Context(), tid, body.ToDomain())
	if err != nil {
		return err
	}

	return c.JSON(presenter.NewTaskUpdateResponse(updatedTask))
}

func (h *TaskHandler) DeleteTask(c fiber.Ctx) error {
	tid := fiber.Params[int](c, "tid")

	if err := h.taskService.DeleteTask(c.Context(), tid); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
