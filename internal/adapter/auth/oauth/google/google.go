package google

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"samsamoohooh-go-api/internal/infra/config"
)

const (
	ScopeEmail   = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile = "https://www.googleapis.com/auth/userinfo.profile"
)

type Google struct {
	oauth2Config oauth2.Config
	config       *config.Config
}

func New(config *config.Config) *Google {
	return &Google{config: config, oauth2Config: oauth2.Config{
		ClientID:     config.Oauth.Google.ClientID,
		ClientSecret: config.Oauth.Google.ClientSecret,
		RedirectURL:  config.Oauth.Google.CallbackURL,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}}
}

func (g *Google) GetLoginURL(state string) string {
	return g.oauth2Config.AuthCodeURL(state)
}

func (g *Google) GetSub(ctx context.Context, code string) (string, error) {
	token, err := g.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return "", err
	}

	client := g.oauth2Config.Client(ctx, token)
	resp, err := client.Get(g.config.Oauth.Google.UserInfoURL)
	if err != nil {
		return "", err
	}

	userInfoResp := new(UserInfoResponse)
	err = json.NewDecoder(resp.Body).Decode(userInfoResp)
	if err != nil {
		return "", err
	}

	return userInfoResp.Sub, nil
}
