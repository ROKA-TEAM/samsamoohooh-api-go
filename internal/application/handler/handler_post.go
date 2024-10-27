package handler

import (
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
)

type PostHandler struct {
	postService port.PostService
}

func NewPostHandler(postService port.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (h *PostHandler) Route(router fiber.Router) {
	router.Post("/", h.CreatePost)
	router.Get("/:id/comments", h.GetCommentsByPostID)
	router.Put("/:id", h.UpdatePost)
	router.Delete("/:id", h.DeletePost)
}

func (h *PostHandler) CreatePost(c fiber.Ctx) error {
	req := new(presenter.PostCreateRequest)

	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	createdPost, err := h.postService.CreatePost(c.Context(), req.GroupID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(presenter.NewPostCreateResponse(createdPost))
}

func (h *PostHandler) GetCommentsByPostID(c fiber.Ctx) error {
	req := new(presenter.PostGetCommentsByPostIDRequest)

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	if err := c.Bind().Query(req); err != nil {
		return err
	}

	gotComments, err := h.postService.GetCommentsByPostID(c.Context(), req.ID, req.Limit, req.Offset)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewPostGetCommentsByIDResponse(gotComments))
}

func (h *PostHandler) UpdatePost(c fiber.Ctx) error {
	req := new(presenter.PostUpdateRequest)

	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	updatedPost, err := h.postService.UpdatePost(c.Context(), req.ID, req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewPostUpdateResponse(updatedPost))
}

func (h *PostHandler) DeletePost(c fiber.Ctx) error {
	req := new(presenter.PostDeleteRequest)

	if err := c.Bind().URI(req); err != nil {
		return err
	}

	err := h.postService.DeletePost(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
