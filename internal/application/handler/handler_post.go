package handler

import (
	domain "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
)

type PostHandler struct {
	postService domain.PostService
}

func NewPostHandler(
	postService domain.PostService,
) *PostHandler {
	return &PostHandler{postService: postService}
}

func (h *PostHandler) CreatePost(c fiber.Ctx) error {
	body := new(presenter.PostCreateRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	createdPost, err := h.postService.CreatePost(c.Context(), body.GroupID, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewPostCreateResponse(createdPost))
}

func (h *PostHandler) GetCommentsByPostID(c fiber.Ctx) error {
	pid := fiber.Params[int](c, "pid")
	limit := fiber.Query[int](c, "limit", DefaultLimit)
	offset := fiber.Query[int](c, "offset", DefaultOffset)

	listComment, err := h.postService.GetCommentsByPostID(c.Context(), pid, offset, limit)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewPostGetCommentsByIDResponse(listComment))
}

func (h *PostHandler) UpdatePost(c fiber.Ctx) error {
	pid := fiber.Params[int](c, "pid")
	body := new(presenter.PostUpdateRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	updatedPost, err := h.postService.UpdatePost(c.Context(), pid, body.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewPostUpdateResponse(updatedPost))
}

func (h *PostHandler) DeletePost(c fiber.Ctx) error {
	pid := fiber.Params[int](c, "pid")

	err := h.postService.DeletePost(c.Context(), pid)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
