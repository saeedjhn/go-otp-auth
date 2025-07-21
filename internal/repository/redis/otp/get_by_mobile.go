package otp

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/saeedjhn/go-otp-auth/pkg/richerror"
)

func (d DB) GetByMobile(ctx context.Context, mobile string) (string, error) {
	value, err := d.conn.Client().Get(ctx, mobile).Result()
	if err != nil {
		rErr := redis.Nil
		if errors.As(err, &rErr) {
			return "", nil
		}

		return "", richerror.New(_opRedisOTPGetByMobile).
			WithErr(err).
			WithKind(richerror.KindStatusInternalServerError)
	}

	return value, nil
}
