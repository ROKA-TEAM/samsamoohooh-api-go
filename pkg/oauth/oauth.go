package oauth

import "context"

type Payload struct {
	Name string
	Sub  string
}

type AuthorizationGrantCodeService interface {
	GetLoginURL(state string) string
	Exchange(ctx context.Context, code string) (*Payload, error)
	AuthenticateOrRegister(ctx context.Context, code string) (string, string, error)
}
