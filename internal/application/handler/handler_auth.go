package handler

import (
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/presenter"
	"samsamoohooh-go-api/internal/infra/exception"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	tokenService port.TokenService
}

func NewAuthHandler(tokenService port.TokenService) *AuthHandler {
	return &AuthHandler{
		tokenService: tokenService,
	}
}

func (h *AuthHandler) Route(router fiber.Router) {
	router.Post("/token/refresh", h.Refresh)
	router.Post("/token/validation", h.Validation)
}

func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	req := new(presenter.AuthRefreshRequest)
	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	parsedToken, err := h.tokenService.ParseToken(req.RefreshToken)
	if err != nil {
		return err
	}

	accessToken, err := h.tokenService.GenerateAccessTokenString(parsedToken.ID, parsedToken.Role)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewAuthRefreshResponse(accessToken))

}

func (h *AuthHandler) Validation(c fiber.Ctx) error {
	req := new(presenter.AuthValidationRequest)
	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	if req.AccessToken == "" && req.RefreshToken == "" {
		return exception.NewWithoutErr(
			exception.ErrValidation,
			exception.StatusBadRequest,
			"accessToken or refreshToken is required",
		)
	}

	if req.AccessToken != "" {
		_, err := h.tokenService.ParseToken(req.AccessToken)
		if err != nil {
			return err
		}
	}

	if req.RefreshToken != "" {
		_, err := h.tokenService.ParseToken(req.RefreshToken)
		if err != nil {
			return err
		}
	}

	return c.SendStatus(fiber.StatusNoContent)
}
