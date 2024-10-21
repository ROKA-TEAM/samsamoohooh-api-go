package google

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/repository/database/ent"
	"samsamoohooh-go-api/pkg/oauth"
	"samsamoohooh-go-api/pkg/token"
)

const (
	scopeProfile = "https://www.googleapis.com/auth/userinfo.profile"
	userInfoAPI  = "https://www.googleapis.com/oauth2/v3/userinfo"
)

var _ oauth.AuthorizationGrantCodeService = (*Service)(nil)

type Service struct {
	config       *config.Config
	oauth2       *oauth2.Config
	userService  domain.UserService
	tokenService token.Service
}

func NewService(
	tokenService token.Service,
	userService domain.UserService,
	config *config.Config,
) *Service {
	return &Service{
		tokenService: tokenService,
		userService:  userService,
		config:       config,
		oauth2: &oauth2.Config{
			ClientID:     config.Oauth.Google.ClientID,
			ClientSecret: config.Oauth.Google.ClientSecret,
			RedirectURL:  config.Oauth.Google.CallbackURL,
			Scopes:       []string{scopeProfile},
			Endpoint:     google.Endpoint,
		},
	}
}

func (s Service) GetLoginURL(state string) string {
	return s.oauth2.AuthCodeURL(state)
}

func (s Service) Exchange(ctx context.Context, code string) (*oauth.Payload, error) {
	t, err := s.oauth2.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := s.oauth2.Client(ctx, t)
	resp, err := client.Get(userInfoAPI)
	if err != nil {
		return nil, err
	}

	var respBody exchangeResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return respBody.toDomain(), nil
}

func (s Service) AuthenticateOrRegister(ctx context.Context, code string) (string, string, error) {
	payload, err := s.Exchange(ctx, code)
	if err != nil {
		return "", "", err
	}

	user, err := s.userService.GetByUserSub(ctx, payload.Sub)
	if ent.IsNotFound(err) {
		createdUser, err := s.userService.CreateUser(ctx, &domain.User{
			Name:      payload.Name,
			Role:      domain.UserRoleUser,
			Social:    domain.UserSocialGoogle,
			SocialSub: payload.Sub,
		})
		if err != nil {
			return "", "", err
		}
		user = createdUser
	} else if err != nil {
		return "", "", err
	}

	// 토큰을 발급한다.
	accessToken, err := s.tokenService.GenerateAccessTokenString(user.ID, string(user.Role))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := s.tokenService.GenerateRefreshTokenString(user.ID, string(user.Role))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}
