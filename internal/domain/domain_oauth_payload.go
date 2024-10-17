package domain

import "context"

type OauthPayload struct {
	Name string
	Sub  string
}

type OauthAuthorizationGrantService interface {
	GetLoginURL(state string) string
	Exchange(ctx context.Context, code string) (*OauthPayload, error)
	AuthenticateOrRegister(ctx context.Context, code string) (string, string, error)
}
