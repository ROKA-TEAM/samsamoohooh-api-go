package port

import "context"

type OauthService interface {
	GetLoginURL(state string) string
	GetSub(ctx context.Context, code string) (string, error)
}
