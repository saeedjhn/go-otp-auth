package bootstrap

import (
	"github.com/saeedjhn/go-otp-auth/configs"
	"github.com/saeedjhn/go-otp-auth/internal/adapter/jsonfilelogger"
	"github.com/saeedjhn/go-otp-auth/internal/contract"
)

func NewLogger(configApp configs.Application, configLogger jsonfilelogger.Config) contract.Logger {
	strategy := createEnvironmentStrategy(configApp.Env, configLogger)

	return jsonfilelogger.New(strategy).Configure()
}

func createEnvironmentStrategy(env configs.Env, config jsonfilelogger.Config) jsonfilelogger.EnvironmentStrategy {
	if env == configs.Production {
		return jsonfilelogger.NewProductionStrategy(config)
	}

	return jsonfilelogger.NewDevelopmentStrategy(config)
}
