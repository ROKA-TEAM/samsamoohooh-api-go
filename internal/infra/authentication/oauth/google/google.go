package google

import (
	"context"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/exception"

	"github.com/carlmjohnson/requests"
)

var _ port.ImplictGrantService = (*GoogleOauthService)(nil)

type GoogleOauthService struct {
	config         *config.Config
	userRepository port.UserRepository
	tokenService   port.TokenService
}

func NewGoogleOauthService(
	config *config.Config,
	userRepository port.UserRepository,
	tokenService port.TokenService,
) *GoogleOauthService {
	return &GoogleOauthService{
		config:         config,
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (s *GoogleOauthService) Authenticate(ctx context.Context, tokenStrig string) (string, string, error) {
	responseBody := struct {
		Sub  string `json:"sub"`
		Name string `json:"name"`
	}{}

	err := requests.
		URL(s.config.Oauth.Google.GetUserInfoURL).
		Header("Authorization", "Bearer "+tokenStrig).
		ToJSON(&responseBody).
		Fetch(ctx)

	if err != nil {
		return "", "", exception.New(
			err,
			exception.ErrOauthRequestFailed,
			exception.StatusUnauthorized,
			"Failed to get user info from Google",
			exception.WithData(
				exception.Map{
					"token": tokenStrig,
				},
			),
		)
	}

	user, err := s.userRepository.GetByUserSub(ctx, responseBody.Sub)
	if exception.Is(err, exception.ErrMySQLNotFound) {
		exception.Is(err, exception.ErrMySQLNotFound)
		newUser := &domain.User{
			SocialSub: responseBody.Sub,
			Social:    domain.UserSocialGoogle,
			Name:      responseBody.Name,
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
