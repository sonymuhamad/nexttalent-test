package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/subosito/gotenv"
)

var Config EnvConfig

type EnvConfig struct {
	APIPort           int    `envconfig:"API_PORT"`
	DatabaseURL       string `envconfig:"DATABASE_URL"`
	HttpClientTimeOut int64  `envconfig:"HTTP_CLIENT_TIME_OUT"`
	TimeApiURL        string `envconfig:"TIME_API_URL"`
}

func LoadEnvConfig() EnvConfig {
	_ = gotenv.Load()

	var cfg EnvConfig

	err := envconfig.Process("", &cfg)
	if err != nil {
		panic(err)
	}

	Config = cfg

	return cfg
}
