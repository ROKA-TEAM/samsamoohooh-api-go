package port

import "context"

type ImplictGrantService interface {
	Authenticate(ctx context.Context, tokenStrig string) (string, string, error)
}
