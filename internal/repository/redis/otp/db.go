package otp

import (
	"github.com/saeedjhn/go-otp-auth/pkg/persistance/cache/redis"
)

type DB struct {
	conn *redis.DB
}

func New(conn *redis.DB) DB {
	return DB{conn: conn}
}
