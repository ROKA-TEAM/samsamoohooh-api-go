package utils

import (
	"samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/pkg/box"

	"github.com/redis/go-redis/v9"
)

func Wrap(err error) error {
	switch {
	case box.Equal(err, redis.Nil):
		return box.Wrap(domain.ErrBadRequest, "value is missing, or has expired")
	case nil == err:
		return nil
	}
	return box.Wrap(domain.ErrInternal, err.Error())
}
