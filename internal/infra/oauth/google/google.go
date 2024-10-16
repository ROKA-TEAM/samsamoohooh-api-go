package google

import (
	"context"
	"encoding/json"
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/infra/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	scopeProfile = "https://www.googleapis.com/auth/userinfo.profile"

	userInfoAPI = "https://www.googleapis.com/oauth2/v3/userinfo"
)

// Authorization Code Grant 방식으로 구현

type OauthGoogleService struct {
	config      *config.Config
	oauthConfig *oauth2.Config
}

func NewOauthGoogleService(config *config.Config) *OauthGoogleService {
	return &OauthGoogleService{
		config: config,
		oauthConfig: &oauth2.Config{
			ClientID:     config.Oauth.Google.ClientID,
			ClientSecret: config.Oauth.Google.ClientSecret,
			RedirectURL:  config.Oauth.Google.CallbackURL,
			Scopes:       []string{scopeProfile},
			Endpoint:     google.Endpoint,
		},
	}
}

func (s OauthGoogleService) GetLoginURL(state string) string {
	return s.oauthConfig.AuthCodeURL(state)
}

func (s OauthGoogleService) Exchange(ctx context.Context, code string) (*domain.OauthPayload, error) {
	token, err := s.oauthConfig.Exchange(ctx, code)
	if err != nil {
		// TODO: 에러 처리
		return nil, err
	}

	client := s.oauthConfig.Client(ctx, token)
	resp, err := client.Get(userInfoAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respBody exchangeResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody.ToDomain(), nil
}
