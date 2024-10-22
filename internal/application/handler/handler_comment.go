package handler

import (
	domain "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
)

type CommentHandler struct {
	commentService domain.CommentService
}

func NewCommentHandler(
	commentService domain.CommentService,
) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) Route(router fiber.Router) {
	router.Post("/", h.CreateComment)
	router.Get("/:cid", h.GetByCommentID)
	router.Put("/:cid", h.UpdateComment)
	router.Delete("/:cid", h.DeleteComment)
}

func (h *CommentHandler) CreateComment(c fiber.Ctx) error {
	body := new(presenter.CommentCreateRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	createdComment, err := h.commentService.CreateComment(c.Context(), body.PostID, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewCommentCreateResponse(createdComment))
}

func (h *CommentHandler) GetByCommentID(c fiber.Ctx) error {
	cid := fiber.Params[int](c, "cid")

	gotComment, err := h.commentService.GetByCommentID(c.Context(), cid)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewCommentGetByCommentIDResponse(gotComment))
}

func (h *CommentHandler) UpdateComment(c fiber.Ctx) error {
	cid := fiber.Params[int](c, "cid")

	body := new(presenter.CommentUpdateRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	updatedComment, err := h.commentService.UpdateComment(c.Context(), cid, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewCommentUpdateResponse(updatedComment))
}

func (h *CommentHandler) DeleteComment(c fiber.Ctx) error {
	cid := fiber.Params[int](c, "cid")

	err := h.commentService.DeleteComment(c.Context(), cid)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
