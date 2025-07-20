package configs

import (
	"time"

	"github.com/saeedjhn/go-otp-auth/internal/adapter/jsonfilelogger"

	"github.com/saeedjhn/go-otp-auth/internal/service/authentication"
	"github.com/saeedjhn/go-otp-auth/pkg/persistance/cache/redis"
	"github.com/saeedjhn/go-otp-auth/pkg/persistance/db/mysql"
)

type Env string

const (
	Development Env = "development"
	Production  Env = "production"
)

func (e Env) String() string {
	return string(e)
}

type Application struct {
	Name                    string        `mapstructure:"name"`
	Version                 string        `mapstructure:"version"`
	Debug                   bool          `mapstructure:"debug"`
	Env                     Env           `mapstructure:"env"`
	EntropyPassword         float64       `mapstructure:"entropy_password"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
}

type HTTPServer struct {
	Host    string        `mapstructure:"host"`
	Port    string        `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Config struct {
	Application Application           `mapstructure:"application"`
	HTTPServer  HTTPServer            `mapstructure:"http_server"`
	Logger      jsonfilelogger.Config `mapstructure:"logger"`
	Mysql       mysql.Config          `mapstructure:"mysql"`
	Redis       redis.Config          `mapstructure:"redis"`
	Auth        authentication.Config `mapstructure:"auth"`
}
