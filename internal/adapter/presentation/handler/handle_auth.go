package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"samsamoohooh-go-api/internal/core/domain"
	"samsamoohooh-go-api/internal/core/dto"
	"samsamoohooh-go-api/internal/core/port"
	"strconv"
	"time"
)

var store = session.New(session.Config{
	Expiration: time.Minute * 3,
})

type AuthHandler struct {
	googleOauthService port.OauthService
	userRepository     port.UserRepository
	jwtService         port.JWTService
	authService        port.AuthService
}

func NewAuthHandler(googleOauthService port.OauthService, userRepository port.UserRepository, jwtService port.JWTService, authService port.AuthService) *AuthHandler {
	return &AuthHandler{googleOauthService: googleOauthService, userRepository: userRepository, jwtService: jwtService, authService: authService}
}

func (h *AuthHandler) MoreInfo(c fiber.Ctx) error {
	//  -> AllowEntryOnlyTempTokenMiddleware -> call this func
	body := new(dto.AuthMoreInfoRequest)

	if err := c.Bind().JSON(body); err != nil {
		return err
	}

	payload := fiber.Locals[domain.TempTokenPayload](c, domain.Temp)
	id, err := strconv.Atoi(payload.Subject)
	if err != nil {
		return errors.Wrap(domain.ErrInternal, err.Error())
	}

	updatedUser, err := h.userRepository.Update(c.Context(), uint(id), body.ToDomain())
	if err != nil {
		return err
	}

	accessTokenString, err := h.jwtService.CreateAccessToken(updatedUser)
	if err != nil {
		return err
	}

	refreshTokenString, err := h.jwtService.CreateRefreshToken(updatedUser)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&dto.AuthMoreInfoResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	})
}

func (h *AuthHandler) GoogleLogin(c fiber.Ctx) error {
	// session 설정
	sess, err := store.Get(c)
	if err != nil {
		return err
	}

	state, err := h.authService.GenerateSecureToken(12)
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

	sess.Delete("state")

	sub, err := h.googleOauthService.GetSub(c.Context(), c.FormValue("code"))
	if err != nil {
		return err
	}

	// 이미 존재하는 sub인지 확인하기 위해, 사용자 조회
	queriedUser, err := h.userRepository.GetBySub(c.Context(), sub)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 사용자 만들기
		createdUser, err := h.userRepository.Create(c.Context(), &domain.User{
			Role:   domain.Guest,
			Social: domain.Google,
			Sub:    sub,
		})
		if err != nil {
			return err
		}

		// temporary token 만들기
		tokenString, err := h.jwtService.CreateTempToken(createdUser.ID, createdUser.Sub, string(createdUser.Social))
		if err != nil {
			return err
		}

		// 임시 토큰 발급해 반환해주기
		return c.Status(fiber.StatusOK).JSON(&dto.AuthGoogleCallbackResponse{
			TempToken: tokenString,
		})

	} else if err != nil {
		return err
	}

	// 만약 query한 사용자의 정보 중 name와 Resolution이 비워져 있다면
	if len(queriedUser.Name) == 0 || len(queriedUser.Resolution) == 0 {
		// 	임시토큰 발급
		tokenString, err := h.jwtService.CreateTempToken(queriedUser.ID, queriedUser.Sub, string(queriedUser.Social))
		if err != nil {
			return err
		}

		// 임시 토큰 발급해 반환해주기
		return c.Status(fiber.StatusOK).JSON(&dto.AuthGoogleCallbackResponse{
			TempToken: tokenString,
		})

	}

	// 아니라면 정식 토큰 발급

	accessTokenString, err := h.jwtService.CreateAccessToken(queriedUser)
	if err != nil {
		return err
	}

	refreshTokenString, err := h.jwtService.CreateRefreshToken(queriedUser)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(&dto.AuthGoogleCallbackResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	})
}
