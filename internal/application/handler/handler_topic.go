package handler

import (
	domain "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
)

type TopicHandler struct {
	topicService domain.TopicService
}

func NewTopicHandler(
	topicService domain.TopicService,
) *TopicHandler {
	return &TopicHandler{topicService: topicService}
}

func (h *TopicHandler) Route(router fiber.Router) {
	router.Post("/", h.CreateTopic)
	router.Get("/:tid", h.GetByTopicID)
	router.Put("/:tid", h.UpdateTopic)
	router.Delete("/:tid", h.Delete)
}

func (h *TopicHandler) CreateTopic(c fiber.Ctx) error {
	body := new(presenter.TopicCreateRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	createdTopic, err := h.topicService.CreateTopic(c.Context(), body.TaskID, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewTopicCreateResponse(createdTopic))
}

func (h *TopicHandler) GetByTopicID(c fiber.Ctx) error {
	tid := fiber.Params[int](c, "tid")

	gotTopic, err := h.topicService.GetByTopicID(c.Context(), tid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTopicGetByIDResponse(gotTopic))
}

func (h *TopicHandler) UpdateTopic(c fiber.Ctx) error {
	tid := fiber.Params[int](c, "tid")

	body := new(presenter.TopicUpdateRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	updatedTopic, err := h.topicService.UpdateTopic(c.Context(), tid, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTopicUpdateResponse(updatedTopic))

}

func (h *TopicHandler) Delete(c fiber.Ctx) error {
	tid := fiber.Params[int](c, "tid")

	if err := h.topicService.DeleteTopic(c.Context(), tid); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
