package handler

import (
	domain2 "samsamoohooh-go-api/internal/application/domain"
)

type CommentHandler struct {
	commentService domain2.CommentService
}

//
//func NewCommentHandler(commentService domain2.CommentService) *CommentHandler {
//	return &CommentHandler{commentService: commentService}
//}
//
//func (h *CommentHandler) Route(r fiber.Router, guard *middleware.GuardMiddleware) {
//	r.Post("/", guard.RequireAccess(domain2.UserRoleAdmin), h.Create)
//	r.Get("/", h.List, guard.RequireAccess(domain2.UserRoleAdmin))
//	r.Get("/:id", h.GetByID, guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest))
//	r.Put("/:id", h.Update, guard.RequireAccess(domain2.UserRoleAdmin, domain2.UserRoleGuest))
//	r.Delete("/:id", h.Delete, guard.RequireAccess(domain2.UserRoleAdmin))
//}
//
//func (h *CommentHandler) Create(c *fiber.Ctx) error {
//	body := new(presenter.CommentCreateRequest)
//	if err := utils.ParseAndVerify(c, body); err != nil {
//		return err
//	}
//
//	createdComment, err := h.commentService.Create(c.Context(), body.PostID, body.ToDomain())
//	if err != nil {
//		return err
//	}
//
//	return c.Status(fiber.StatusCreated).JSON(presenter.NewCommentCreateResponse(createdComment))
//}
//
//func (h *CommentHandler) List(c *fiber.Ctx) error {
//	limit := c.QueryInt("limit", DefaultLimit)
//	offset := c.QueryInt("offset", DefaultOffset)
//
//	listComment, err := h.commentService.List(c.Context(), limit, offset)
//	if err != nil {
//		return err
//	}
//
//	return c.Status(fiber.StatusOK).JSON(presenter.NewCommentListResponse(listComment))
//}
//
//func (h *CommentHandler) GetByID(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	if err != nil {
//		return err
//	}
//
//	comment, err := h.commentService.GetByID(c.Context(), id)
//	if err != nil {
//		return err
//	}
//
//	return c.Status(fiber.StatusOK).JSON(presenter.NewCommentGetByIDResponse(comment))
//}
//
//func (h *CommentHandler) Update(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	if err != nil {
//		return err
//	}
//
//	body := new(presenter.CommentUpdateRequest)
//	if err := utils.ParseAndVerify(c, body); err != nil {
//		return err
//	}
//
//	updatedComment, err := h.commentService.Update(c.Context(), id, body.ToDomain())
//	if err != nil {
//		return err
//	}
//
//	return c.Status(fiber.StatusOK).JSON(presenter.NewCommentUpdateResponse(updatedComment))
//}
//
//func (h *CommentHandler) Delete(c *fiber.Ctx) error {
//	id, err := c.ParamsInt("id")
//	if err != nil {
//		return err
//	}
//
//	err = h.commentService.Delete(c.Context(), id)
//	if err != nil {
//		return err
//	}
//
//	return c.SendStatus(fiber.StatusNoContent)
//}
