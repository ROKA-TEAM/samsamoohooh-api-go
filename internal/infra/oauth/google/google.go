package google

import (
	"context"
	"encoding/json"
	domain2 "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/infra/config"

	"github.com/pkg/errors"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var _ domain2.OauthAuthorizationGrantService = (*OauthGoogleService)(nil)

const (
	scopeProfile = "https://www.googleapis.com/auth/userinfo.profile"

	userInfoAPI = "https://www.googleapis.com/oauth2/v3/userinfo"
)

// Authorization Code Grant 방식으로 구현

type OauthGoogleService struct {
	config       *config.Config
	oauthConfig  *oauth2.Config
	userService  domain2.UserService
	tokenService domain2.TokenService
}

func NewOauthGoogleService(
	config *config.Config,
	userService domain2.UserService,
	tokenService domain2.TokenService,
) *OauthGoogleService {
	return &OauthGoogleService{
		config: config,
		oauthConfig: &oauth2.Config{
			ClientID:     config.Oauth.Google.ClientID,
			ClientSecret: config.Oauth.Google.ClientSecret,
			RedirectURL:  config.Oauth.Google.CallbackURL,
			Scopes:       []string{scopeProfile},
			Endpoint:     google.Endpoint,
		},
		userService:  userService,
		tokenService: tokenService,
	}
}

func (s OauthGoogleService) GetLoginURL(state string) string {
	return s.oauthConfig.AuthCodeURL(state)
}

func (s OauthGoogleService) Exchange(ctx context.Context, code string) (*domain2.OauthPayload, error) {
	token, err := s.oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, errors.Wrap(domain2.ErrInternal, err.Error())
	}

	client := s.oauthConfig.Client(ctx, token)
	resp, err := client.Get(userInfoAPI)
	if err != nil {
		return nil, errors.Wrap(domain2.ErrInternal, err.Error())
	}

	var respBody exchangeResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, errors.Wrap(domain2.ErrInternal, err.Error())
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, errors.Wrap(domain2.ErrInternal, err.Error())
	}

	return respBody.toDomain(), nil
}

func (s OauthGoogleService) AuthenticateOrRegister(ctx context.Context, code string) (string, string, error) {
	payload, err := s.Exchange(ctx, code)
	if err != nil {
		return "", "", errors.Wrap(domain2.ErrInternal, err.Error())
	}

	user, err := s.userService.GetBySub(ctx, payload.Sub)

	if errors.Is(err, domain2.ErrNotFound) {
		createdUser, err := s.userService.Create(ctx, &domain2.User{
			Name:      payload.Name,
			Role:      domain2.UserRoleGuest,
			Social:    domain2.UserSocialGoogle,
			SocialSub: payload.Sub,
		})
		if err != nil {
			return "", "", err
		}

		user = createdUser
	} else if err != nil {
		return "", "", err
	}
	// 전에 등록한 사용자이다.

	// 토큰을 발급한다.
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
