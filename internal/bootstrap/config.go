package bootstrap

import "github.com/saeedjhn/go-otp-auth/configs"

func ConfigLoad(option configs.Option) (*configs.Config, error) {
	return configs.Load(option)
}
