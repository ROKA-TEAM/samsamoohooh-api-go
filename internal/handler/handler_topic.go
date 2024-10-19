package handler

import (
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/handler/utils"
	"samsamoohooh-go-api/internal/infra/middleware"
	"samsamoohooh-go-api/internal/infra/presenter"

	"github.com/gofiber/fiber/v2"
)

type TopicHandler struct {
	topicService domain.TopicService
}

func NewTopicHandler(topicService domain.TopicService) *TopicHandler {
	return &TopicHandler{topicService: topicService}
}

func (h *TopicHandler) Route(router fiber.Router, guard *middleware.GuardMiddleware) {
	router.Post("/", guard.RequireAccess(domain.UserRoleAdmin, domain.UserRoleGuest), h.Create)
	router.Get("/", guard.RequireAccess(domain.UserRoleAdmin), h.List)
	router.Get("/:id", guard.RequireAccess(domain.UserRoleAdmin, domain.UserRoleGuest), h.GetByID)
	router.Put("/:id", guard.RequireAccess(domain.UserRoleAdmin, domain.UserRoleGuest), h.Update)
	router.Delete("/:id", guard.RequireAccess(domain.UserRoleAdmin), h.Delete)
}

func (h *TopicHandler) Create(c *fiber.Ctx) error {
	body := new(presenter.TopicCreateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	createdTopic, err := h.topicService.Create(c.Context(), body.TaskID, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewTopicCreateResponse(createdTopic))
}

func (h *TopicHandler) List(c *fiber.Ctx) error {
	offset := c.QueryInt("offset", DefaultOffset)
	limit := c.QueryInt("limit", DefaultLimit)

	listTopics, err := h.topicService.List(c.Context(), offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTopicListResponse(listTopics))
}

func (h *TopicHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	gotTopic, err := h.topicService.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTopicGetByIDResponse(gotTopic))
}

func (h *TopicHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	body := new(presenter.TopicUpdateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	updatedTopic, err := h.topicService.Update(c.Context(), id, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewTopicUpdateResponse(updatedTopic))
}

func (h *TopicHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = h.topicService.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
