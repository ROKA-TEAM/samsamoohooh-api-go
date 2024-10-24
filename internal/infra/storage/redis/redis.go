package redis

import (
	"context"
	"samsamoohooh-go-api/internal/application/port"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/internal/infra/exception"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ port.RedisRepository = (*Redis)(nil)

type Redis struct {
	client *redis.Client
}

func NewRdis(config *config.Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &Redis{client}, nil
}

func (r *Redis) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	err := r.client.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return exception.New(
			err,
			exception.ErrRedisSetFailed,
			exception.StatusInternalServerError,
			"failed to set redis",
			exception.WithData(
				exception.Map{
					"key":   key,
					"value": value,
				},
			),
		)
	}

	return nil
}

func (r *Redis) GetBytes(ctx context.Context, key string) ([]byte, error) {
	res, err := r.client.Get(ctx, key).Result()

	if err != nil {
		return nil, exception.New(
			err,
			exception.ErrRedisGetFailed,
			exception.StatusInternalServerError,
			"failed to get redis",
			exception.WithData(
				exception.Map{
					"key": key,
				},
			),
		)
	}

	bytes := []byte(res)
	return bytes, nil
}

func (r *Redis) GetInt(ctx context.Context, key string) (int, error) {
	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return 0, exception.New(
			err,
			exception.ErrRedisGetFailed,
			exception.StatusInternalServerError,
			"failed to get redis",
			exception.WithData(
				exception.Map{
					"key": key,
				},
			),
		)
	}

	return strconv.Atoi(res)
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return exception.New(
			err,
			exception.ErrRedisDeleteFailed,
			exception.StatusInternalServerError,
			"failed to delete redis",
			exception.WithData(
				exception.Map{
					"key": key,
				},
			),
		)
	}

	return nil
}
