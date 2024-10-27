package handler

import (
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
)

type CommentHandler struct {
	commentService port.CommentService
}

func NewCommentHandler(commentService port.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

func (h *CommentHandler) Route(router fiber.Router) {
	router.Post("", h.CreatePost)
	router.Put("/:id", h.UpdatePost)
	router.Delete("/:id", h.DeletePost)
}

func (h *CommentHandler) CreatePost(c fiber.Ctx) error {
	req := new(presenter.CommentCreateRequest)

	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	createdComment, err := h.commentService.CreateComment(c.Context(), req.PostID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewCommentCreateResponse(createdComment))
}

func (h *CommentHandler) UpdatePost(c fiber.Ctx) error {
	req := new(presenter.CommentUpdateRequest)

	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	updatdComment, err := h.commentService.UpdateComment(c.Context(), req.ID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewCommentUpdateResponse(updatdComment))
}

func (h *CommentHandler) DeletePost(c fiber.Ctx) error {
	req := new(presenter.CommentDeleteRequest)

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	err := h.commentService.DeleteComment(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
