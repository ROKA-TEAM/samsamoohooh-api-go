package handler

import (
	domain "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/handler/utils"
	"samsamoohooh-go-api/internal/application/presenter"
	"samsamoohooh-go-api/pkg/box"
	"samsamoohooh-go-api/pkg/oauth"
	"samsamoohooh-go-api/pkg/token"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
)

var store = session.New(session.Config{
	Expiration: time.Minute * 3,
})

const (
	stateKey = "state"
)

type AuthHandler struct {
	kakaoOauthService  oauth.AuthorizationGrantCodeService
	googleOauthService oauth.AuthorizationGrantCodeService
	tokenService       token.Service
}

func NewAuthHandler(
	kakaoOauthService oauth.AuthorizationGrantCodeService,
	googleOauthService oauth.AuthorizationGrantCodeService,
	tokenService token.Service) *AuthHandler {
	return &AuthHandler{
		kakaoOauthService:  kakaoOauthService,
		googleOauthService: googleOauthService,
		tokenService:       tokenService,
	}
}
func (h *AuthHandler) Route(router fiber.Router) {
	token := router.Group("/token")
	{
		token.Post("/refresh", h.Refresh)
		token.Post("/validation", h.Validation)
	}
	router.Get("/google", h.GetLoginURLOfGoogle)
	router.Get("/google/callback", h.GoogleCallback)
	router.Get("/kakao", h.GetLoginURLOfKakao)
	router.Get("/kakao/callback", h.KaKaoCallback)
}

func (h *AuthHandler) Validation(c fiber.Ctx) error {
	body := new(presenter.AuthValidationRequest)

	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	if body.AccessToken == "" && body.RefreshToken == "" {
		return box.Wrap(domain.ErrAuthorization, "request body is empty")
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

func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	body := new(presenter.AuthRefreshRequest)
	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	_, err := h.tokenService.ValidateToken(body.RefreshToken)
	if err != nil {
		return err
	}

	t, err := h.tokenService.ParseToken(body.RefreshToken)
	if err != nil {
		return err
	}

	accessToken, err := h.tokenService.GenerateAccessTokenString(t.Subject, t.Role)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&presenter.AuthRefreshResponse{
		AccessToken: accessToken,
	})
}

func (h *AuthHandler) GetLoginURLOfGoogle(c fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := utils.GenerateState()
	sess.Set(stateKey, state)
	err = sess.Save()
	if err != nil {
		return err
	}

	redirectURL := h.googleOauthService.GetLoginURL(state)
	return c.Redirect().Status(fiber.StatusTemporaryRedirect).To(redirectURL)
}

func (h *AuthHandler) GoogleCallback(c fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := sess.Get(stateKey)
	sess.Delete(stateKey)
	err = sess.Save()
	if err != nil {
		return err
	}

	if state != c.FormValue(stateKey) {
		return box.Wrap(domain.ErrBadRequest, "invalid state")
	}

	accessToken, refreshToken, err := h.googleOauthService.AuthenticateOrRegister(c.Context(), c.FormValue("code"))
	if err != nil {
		return box.Wrap(domain.ErrBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&presenter.GoogleCallbackResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *AuthHandler) GetLoginURLOfKakao(c fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := utils.GenerateState()
	sess.Set(stateKey, state)
	err = sess.Save()
	if err != nil {
		return err
	}

	redirectURL := h.kakaoOauthService.GetLoginURL(state)
	return c.Redirect().Status(fiber.StatusTemporaryRedirect).To(redirectURL)
}

func (h *AuthHandler) KaKaoCallback(c fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state := sess.Get(stateKey)
	sess.Delete(stateKey)
	err = sess.Save()
	if err != nil {
		return err
	}

	if state != c.FormValue(stateKey) {
		return box.Wrap(domain.ErrBadRequest, "invalid state")
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
