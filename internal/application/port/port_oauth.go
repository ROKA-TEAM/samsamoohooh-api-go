package port

import "context"

type OauthImplictGrantService interface {
	Authenticate(ctx context.Context, tokenStrig string) (string, string, error)
}
