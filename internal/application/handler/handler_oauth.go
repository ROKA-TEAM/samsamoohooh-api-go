package handler

import (
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
)

type OauthHandler struct {
	kakaoService  port.OauthImplictGrantService
	googleService port.OauthImplictGrantService
}

func NewOauthHandler(
	kakaoService port.OauthImplictGrantService,
	googleService port.OauthImplictGrantService,
) *OauthHandler {
	return &OauthHandler{
		kakaoService:  kakaoService,
		googleService: googleService,
	}
}

func (h *OauthHandler) Route(c fiber.Router) {
	c.Post("/google/login", h.GoogleLogin)
	c.Post("/kakao/login", h.KakaoLogin)
}

func (h *OauthHandler) GoogleLogin(c fiber.Ctx) error {
	req := new(presenter.OauthGoogleLoginRequest)
	if err := c.Bind().JSON(&req); err != nil {
		return err
	}

	accessToken, refreshToken, err := h.googleService.Authenticate(c.Context(), req.Token)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewOauthGoogleLoginResponse(accessToken, refreshToken))

}

func (h *OauthHandler) KakaoLogin(c fiber.Ctx) error {
	req := new(presenter.OauthKakaoLoginRequest)
	if err := c.Bind().JSON(&req); err != nil {
		return err
	}

	accessToken, refreshToken, err := h.kakaoService.Authenticate(c.Context(), req.Token)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewOauthKakaoLoginResponse(accessToken, refreshToken))
}
