package handler

import (
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/handler/utils"
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
	kakaoOauthService  domain.OauthAuthorizationGrantService
	googleOauthService domain.OauthAuthorizationGrantService
}

func NewAuthHandler(
	googleOauthService domain.OauthAuthorizationGrantService,
	kakaoOauthService domain.OauthAuthorizationGrantService,
) *AuthHandler {
	return &AuthHandler{
		googleOauthService: googleOauthService,
		kakaoOauthService:  kakaoOauthService,
	}
}

func (h *AuthHandler) Route(router fiber.Router) {
	router.Get("/auth/google", h.GetLoginURLOfGoogle)
	router.Get("/auth/google/callback", h.GoogleCallback)

	router.Get("/auth/kakao", h.GetLoginURLOfKakao)
	router.Get("/auth/kakao/callback", h.KaKaoCallback)
}

func (h *AuthHandler) GetLoginURLOfGoogle(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := utils.GenerateState()
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
		return errors.Wrap(domain.ErrNotMatchState, "invalid state")
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

	state := utils.GenerateState()
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
		return errors.Wrap(domain.ErrNotMatchState, "invalid state")
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
