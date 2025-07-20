package bootstrap

import (
	"github.com/saeedjhn/go-otp-auth/configs"
	usermysql "github.com/saeedjhn/go-otp-auth/internal/repository/mysql/user"
	authservice "github.com/saeedjhn/go-otp-auth/internal/service/authentication"
	userservice "github.com/saeedjhn/go-otp-auth/internal/service/user"
	uservalidator "github.com/saeedjhn/go-otp-auth/internal/validator/user"
	"github.com/saeedjhn/go-otp-auth/pkg/persistance/db/mysql"
)

type Service struct {
	AuthSvc authservice.Service
	UserSvc userservice.Service
}

func NewService(config *configs.Config, mysqlDB *mysql.DB) *Service {
	var (
		userRepo = usermysql.New(mysqlDB)
		// userRdsRepo = userredis.New(cache.Redis) // Or userInMemRepo := inmemoryuser.New(cache.InMem)
	)

	var (
		authSvc = authservice.New(config.Auth)
		userVld = uservalidator.New(config.Application.EntropyPassword)
		userSvc = userservice.New(config, authSvc, userVld, userRepo)
	)

	return &Service{
		AuthSvc: authSvc,
		UserSvc: userSvc,
	}
}
