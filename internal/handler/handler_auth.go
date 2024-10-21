package handler

import (
	domain2 "samsamoohooh-go-api/internal/application/domain"
	utils2 "samsamoohooh-go-api/internal/application/handler/utils"
	"samsamoohooh-go-api/internal/infra/presenter"
	"time"

	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New(session.Config{
	Expiration: time.Minute * 3,
})

type AuthHandler struct {
	kakaoOauthService  domain2.OauthAuthorizationGrantService
	googleOauthService domain2.OauthAuthorizationGrantService
	tokenService       domain2.TokenService
}

func NewAuthHandler(
	googleOauthService domain2.OauthAuthorizationGrantService,
	kakaoOauthService domain2.OauthAuthorizationGrantService,
	tokenService domain2.TokenService,
) *AuthHandler {
	return &AuthHandler{
		googleOauthService: googleOauthService,
		kakaoOauthService:  kakaoOauthService,
		tokenService:       tokenService,
	}
}

func (h *AuthHandler) Route(router fiber.Router) {
	router.Post("/token/refresh", h.Refresh)
	router.Post("/token/validation", h.Validation)

	router.Get("/google", h.GetLoginURLOfGoogle)
	router.Get("/google/callback", h.GoogleCallback)

	router.Get("/kakao", h.GetLoginURLOfKakao)
	router.Get("/kakao/callback", h.KaKaoCallback)
}

func (h *AuthHandler) Validation(c *fiber.Ctx) error {
	body := new(presenter.AuthValidationRequest)
	if err := utils2.ParseAndVerify(c, body); err != nil {
		return err
	}

	if body.AccessToken == "" && body.RefreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("request body is empty")
	}

	if body.AccessToken != "" {
		_, err := h.tokenService.ValidateToken(body.AccessToken)
		if err != nil {
			return err
		}
	}

	if body.RefreshToken != "" {
		_, err := h.tokenService.ValidateToken(body.RefreshToken)
		if err != nil {
			return err
		}
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	body := new(presenter.AuthRefreshRequest)
	if err := utils2.ParseAndVerify(c, body); err != nil {
		return err
	}

	_, err := h.tokenService.ValidateToken(body.RefreshToken)
	if err != nil {
		return err
	}

	token, err := h.tokenService.ParseToken(body.RefreshToken)
	if err != nil {
		return err
	}

	accessToken, err := h.tokenService.GenerateAccessTokenString(token.Subject, token.Role)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&presenter.AuthRefreshResponse{
		AccessToken: accessToken,
	})
}

func (h *AuthHandler) GetLoginURLOfGoogle(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := utils2.GenerateState()
	sess.Set("state", state)
	err = sess.Save()
	if err != nil {
		return err
	}

	redirectURL := h.googleOauthService.GetLoginURL(state)
	return c.Redirect(redirectURL, fiber.StatusTemporaryRedirect)
}

func (h *AuthHandler) GoogleCallback(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := sess.Get("state")
	sess.Delete("state")
	err = sess.Save()
	if err != nil {
		return err
	}

	if state != c.FormValue("state") {
		return errors.Wrap(domain2.ErrNotMatchState, "invalid state")
	}

	accessToken, refreshToken, err := h.googleOauthService.AuthenticateOrRegister(c.Context(), c.FormValue("code"))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&presenter.GoogleCallbackResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *AuthHandler) GetLoginURLOfKakao(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := utils2.GenerateState()
	sess.Set("state", state)
	err = sess.Save()
	if err != nil {
		return err
	}

	redirectURL := h.kakaoOauthService.GetLoginURL(state)
	return c.Redirect(redirectURL, fiber.StatusTemporaryRedirect)
}

func (h *AuthHandler) KaKaoCallback(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := sess.Get("state")
	sess.Delete("state")
	err = sess.Save()
	if err != nil {
		return err
	}

	if state != c.FormValue("state") {
		return errors.Wrap(domain2.ErrNotMatchState, "invalid state")
	}

	accessToken, refreshToken, err := h.kakaoOauthService.AuthenticateOrRegister(c.Context(), c.FormValue("code"))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&presenter.KaKaoCallbackResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})

}
