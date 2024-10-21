package handler

import (
	"github.com/gofiber/fiber/v2"
	domain2 "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/handler/utils"
	"samsamoohooh-go-api/internal/application/presenter/v1"
	"samsamoohooh-go-api/internal/infra/middleware"
)

type CommentHandler struct {
	commentService domain2.CommentService
}

func NewCommentHandler(commentService domain2.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) Route(r fiber.Router, guard *middleware.GuardMiddleware) {
	r.Post("/", guard.RequireAccess(domain2.UserRoleAdmin), h.Create)
	r.Get("/", h.List, guard.RequireAccess(domain2.UserRoleAdmin))
	r.Get("/:id", h.GetByID, guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest))
	r.Put("/:id", h.Update, guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest))
	r.Delete("/:id", h.Delete, guard.RequireAccess(domain2.UserRoleAdmin))
}

func (h *CommentHandler) Create(c *fiber.Ctx) error {
	body := new(v1.CommentCreateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	createdComment, err := h.commentService.Create(c.Context(), body.PostID, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(v1.NewCommentCreateResponse(createdComment))
}

func (h *CommentHandler) List(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", DefaultLimit)
	offset := c.QueryInt("offset", DefaultOffset)

	listComment, err := h.commentService.List(c.Context(), limit, offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(v1.NewCommentListResponse(listComment))
}

func (h *CommentHandler) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	comment, err := h.commentService.GetByID(c.Context(), id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(v1.NewCommentGetByIDResponse(comment))
}

func (h *CommentHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	body := new(v1.CommentUpdateRequest)
	if err := utils.ParseAndVerify(c, body); err != nil {
		return err
	}

	updatedComment, err := h.commentService.Update(c.Context(), id, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(v1.NewCommentUpdateResponse(updatedComment))
}

func (h *CommentHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}

	err = h.commentService.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
