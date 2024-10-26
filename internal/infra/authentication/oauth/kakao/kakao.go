package kakao

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/exception"
	"strconv"
	"time"

	"github.com/carlmjohnson/requests"
)

var _ port.ImplictGrantService = (*KakaoOauthService)(nil)

type KakaoOauthService struct {
	config         *config.Config
	userRepository port.UserRepository
	tokenService   port.TokenService
}

func NewKakaoOauthService(
	config *config.Config,
	userRepository port.UserRepository,
	tokenService port.TokenService,
) *KakaoOauthService {
	return &KakaoOauthService{
		config:         config,
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (s *KakaoOauthService) Authenticate(ctx context.Context, tokenStrig string) (string, string, error) {
	responseBody := struct {
		ID          int       `json:"id"`
		ConnectedAt time.Time `json:"connected_at"`
		Properties  struct {
			Nickname string `json:"nickname"`
		} `json:"properties"`
	}{}

	err := requests.
		URL(s.config.Oauth.Kakao.GetUserInfoURL).
		Header("Authorization", "Bearer "+tokenStrig).
		ToJSON(&responseBody).
		Fetch(ctx)

	if err != nil {
		return "", "", exception.New(
			err,
			exception.ErrOauthRequestFailed,
			exception.StatusUnauthorized,
			"Failed to get user info from KaKao",
			exception.WithData(
				exception.Map{
					"token": tokenStrig,
				},
			),
		)
	}

	kakaoUserSub := strconv.Itoa(responseBody.ID)

	user, err := s.userRepository.GetByUserSub(ctx, kakaoUserSub)
	if exception.Is(err, exception.ErrMySQLNotFound) {
		exception.Is(err, exception.ErrMySQLNotFound)
		newUser := &domain.User{
			SocialSub: kakaoUserSub,
			Social:    domain.UserSocialKaKao,
			Name:      responseBody.Properties.Nickname,
			Role:      domain.UserRoleUser,
		}

		createdUser, err := s.userRepository.CreateUser(ctx, newUser)
		if err != nil {
			return "", "", err
		}

		user = createdUser
	} else if err != nil {
		return "", "", err
	}

	accessToken, err := s.tokenService.GenerateAccessTokenString(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.tokenService.GenerateRefreshTokenString(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}
