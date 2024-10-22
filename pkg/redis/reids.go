package redis

import (
	"context"
	"samsamoohooh-go-api/internal/infra/config"
	"samsamoohooh-go-api/pkg/redis/utils"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// TODO: Generic으로 개발해도 좋을 듯
// But, generic을 이용해 DX를 높이고자 했는데, interface 검사로 인해서 속도가 저하될 수도 있을 것 같아서 현재는 일단 두 개의 타입만 필요함으로 GetInt, GetBytes 형식으로 개발할 예정
// 그러나 많은 타입의 key value를 저장할 경우 generic으로 변경
type KeyValueStore interface {
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	GetBytes(ctx context.Context, key string) ([]byte, error)
	GetInt(ctx context.Context, key string) (int, error)
	Delete(ctx context.Context, key string) error
	Close() error
}

type Redis struct {
	client *redis.Client
}

func NewRedis(ctx context.Context, config *config.Config) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
		Protocol: config.Redis.Protocol,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, utils.Wrap(err)
	}

	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return utils.Wrap(r.client.Set(ctx, key, value, ttl).Err())
}

func (r *Redis) GetBytes(ctx context.Context, key string) ([]byte, error) {
	res, err := r.client.Get(ctx, key).Result()
	bytes := []byte(res)
	return bytes, utils.Wrap(err)
}

func (r *Redis) GetInt(ctx context.Context, key string) (int, error) {
	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return 0, utils.Wrap(err)
	}

	return strconv.Atoi(res)
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	return utils.Wrap(r.client.Del(ctx, key).Err())
}

func (r *Redis) Close() error {
	return utils.Wrap(r.client.Close())
}
