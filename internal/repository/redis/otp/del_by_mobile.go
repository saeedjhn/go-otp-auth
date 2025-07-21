package otp

import (
	"context"

	"github.com/saeedjhn/go-otp-auth/pkg/richerror"
)

func (d DB) DelByMobile(ctx context.Context, mobile string) (bool, error) {
	success, err := d.conn.Client().Del(ctx, mobile).Result()
	if err != nil {
		return false, richerror.New(_opRedisOTPDelByMobile).
			WithErr(err).
			WithKind(richerror.KindStatusInternalServerError)
	}

	if success != 1 {
		return false, nil
	}

	return true, nil
}
