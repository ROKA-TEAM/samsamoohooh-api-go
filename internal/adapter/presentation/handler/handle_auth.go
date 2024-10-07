package handler

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/pkg/errors"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/port"
	"time"
)

var store = session.New(session.Config{
	Expiration: time.Minute * 3,
})

type AuthHandler struct {
	googleOauthService port.OauthService
}

func NewAuthHandler(googleOauthService port.OauthService) *AuthHandler {
	return &AuthHandler{googleOauthService: googleOauthService}
}

// 인증 (사용자에게 로그인 화면을 보여준다)
func (h *AuthHandler) GoogleLogin(c fiber.Ctx) error {
	// session 설정
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state, err := generateSecureToken(16)
	if err != nil {
		return err
	}

	sess.Set("state", state)
	err = sess.Save()
	if err != nil {
		return err
	}

	redirectURL := h.googleOauthService.GetLoginURL(state)

	return c.Redirect().Status(fiber.StatusTemporaryRedirect).To(redirectURL)
}

// 인가 (사용자의 정보를 가져옴)
func (h *AuthHandler) GoogleCallback(c fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state, ok := sess.Get("state").(string)
	if !ok {
		return errors.Wrap(domain.ErrUnauthorized, "invalid state")
	}

	if state != c.FormValue("state") {
		return errors.Wrap(domain.ErrUnauthorized, "invalid state, not match req state and session state")
	}

	sub, err := h.googleOauthService.GetSub(c.Context(), c.FormValue("code"))
	if err != nil {
		return err
	}

	fmt.Println("sub: ", sub)
	return c.SendString(sub)
}

func generateSecureToken(length int) (string, error) {
	// length가 0보다 작으면 에러 반환
	if length < 1 {
		return "", errors.Wrap(domain.ErrInternal, fmt.Sprintf("invalid token length: %d", length))
	}

	bytes := make([]byte, length)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(bytes), nil
}
