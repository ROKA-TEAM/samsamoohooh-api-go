package port

import (
	"context"
	"time"
)

type RedisRepository interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	GetBytes(ctx context.Context, key string) ([]byte, error)
	GetInt(ctx context.Context, key string) (int, error)
	Delete(ctx context.Context, key string) error
}
