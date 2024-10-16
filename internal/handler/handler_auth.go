package handler

import (
	"errors"
	"fmt"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/handler/utils"
	"samsamoohooh-go-api/internal/infra/oauth/google"
	"samsamoohooh-go-api/internal/infra/presenter"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New(session.Config{
	Expiration: time.Minute * 3,
})

type AuthHandler struct {
	userService  domain.UserService
	tokenService domain.TokenService
	// 직접 의존하는 방식을 이용
	oauthGoogleService *google.OauthGoogleService
}

func NewAuthHandler(
	userService domain.UserService,
	tokenService domain.TokenService,
	oauthGoogleService *google.OauthGoogleService,
) *AuthHandler {
	return &AuthHandler{
		userService:        userService,
		tokenService:       tokenService,
		oauthGoogleService: oauthGoogleService,
	}
}

func (h *AuthHandler) Route(router fiber.Router) {
	router.Get("/auth/google", h.GetLoginURLOfGoogle)
	router.Get("/auth/google/callback", h.GoogleCallback)
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

	redirectURL := h.oauthGoogleService.GetLoginURL(state)
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
		// TODO: errors 패키지를 이용하여 에러 처리
		return errors.New("invalid state")
	}

	payload, err := h.oauthGoogleService.Exchange(c.Context(), c.FormValue("code"))
	if err != nil {
		return err
	}

	// 전에 등록했던 사용자인가
	user, err := h.userService.GetBySub(c.Context(), payload.Sub)

	if errors.Is(err, domain.ErrNotFound) {
		// 전에 등록하지 않은 사용자이다.
		createdUser, err := h.userService.Create(c.Context(), &domain.User{
			Name:      payload.Name,
			Role:      domain.UserRoleGuest,
			Social:    domain.UserSocialGoogle,
			SocialSub: payload.Sub,
		})
		if err != nil {
			return err
		}
		fmt.Println("createdUser: ", createdUser)
		

		user = createdUser

	} else if err != nil {
		return err
	}

	fmt.Println("User: ", user)

	// 전에 등록한 사용자이다.
	// 토큰을 발급한다.
	accessToken, err := h.tokenService.GenerateAccessTokenString(user.ID, domain.TokenRoleType(user.Role))
	if err != nil {
		return err
	}

	refreshToken, err := h.tokenService.GenerateRefreshTokenString(user.ID, domain.TokenRoleType(user.Role))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&presenter.GoogleCallbackResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
