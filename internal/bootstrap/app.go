package bootstrap

import (
	"github.com/saeedjhn/go-otp-auth/internal/contract"

	"github.com/saeedjhn/go-otp-auth/pkg/persistance/cache/redis"

	"github.com/saeedjhn/go-otp-auth/configs"
	"github.com/saeedjhn/go-otp-auth/pkg/persistance/db/mysql"
)

type Application struct {
	Config  *configs.Config
	Logger  contract.Logger
	Redis   *redis.DB
	MySQL   *mysql.DB
	Service *Service
}

func App(config *configs.Config) (*Application, error) {
	a := &Application{Config: config}

	if err := a.setup(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Application) setup() error {
	var err error

	a.Logger = NewLogger(a.Config.Application, a.Config.Logger)

	if a.MySQL, err = NewMysqlConnection(a.Config.Mysql); err != nil {
		return err
	}

	if a.Redis, err = NewRedisClient(a.Config.Redis); err != nil {
		return err
	}

	a.Service = NewService(
		a.Config,
		a.MySQL,
	)

	return nil
}

func (a *Application) CloseMysqlConnection() error {
	return CloseMysqlConnection(a.MySQL)
}

func (a *Application) CloseRedisClientConnection() error {
	return CloseRedisClient(a.Redis)
}
