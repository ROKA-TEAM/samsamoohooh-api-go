package google

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
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
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	client := s.oauthConfig.Client(ctx, token)
	resp, err := client.Get(userInfoAPI)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	var respBody exchangeResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternal, err.Error())
	}

	return respBody.ToDomain(), nil
}
