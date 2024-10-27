package handler

import (
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
)

type TopicHandler struct {
	topicService port.TopicService
}

func NewTopicHandler(topicService port.TopicService) *TopicHandler {
	return &TopicHandler{
		topicService: topicService,
	}
}

func (h *TopicHandler) Route(router fiber.Router) {
	router.Post("/", h.CreateTopic)
	router.Put("/:id", h.UpdateTopic)
	router.Delete("/:id", h.DeleteTopic)
}

func (h *TopicHandler) CreateTopic(c fiber.Ctx) error {
	req := new(presenter.TopicCreateRequest)

	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	createdTopic, err := h.topicService.CreateTopic(c.Context(), req.TaskID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewTopicCreateResponse(createdTopic))
}

func (h *TopicHandler) UpdateTopic(c fiber.Ctx) error {
	req := new(presenter.TopicUpdateRequest)

	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	updatedTopic, err := h.topicService.UpdateTopic(c.Context(), req.ID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTopicUpdateResponse(updatedTopic))
}

func (h *TopicHandler) DeleteTopic(c fiber.Ctx) error {
	req := new(presenter.TopicDeleteRequest)

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	if err := h.topicService.DeleteTopic(c.Context(), req.ID); err != nil {
		return err
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
