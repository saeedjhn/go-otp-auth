package otp

import (
	"context"
	"time"

	"github.com/saeedjhn/go-otp-auth/pkg/richerror"
)

func (d DB) SetByMobile(
	ctx context.Context,
	mobile string,
	code string,
	expireTime time.Duration,
) error {
	err := d.conn.Client().Set(ctx, mobile, code, expireTime).Err()
	if err != nil {
		return richerror.New(_opRedisOTPSetByMobile).
			WithErr(err).
			WithKind(richerror.KindStatusInternalServerError)
	}

	return nil
}
