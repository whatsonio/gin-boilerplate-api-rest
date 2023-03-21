package config

import (
	"app/infra/logger"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Env  string `env:"ENV"`
	Db   DbConfiguration
	Host string `env:"APP_HOST"`
	Port string `env:"APP_PORT"`
	Auth AuthConfiguration
}

func Init() {

	if os.Getenv("ENV") != "PROD" {
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			logger.Fatalf("Error while reading config file %s", err)
		}
	} else {
		viper.AutomaticEnv()
	}

}

func GetConfig() Config {
	return Config{
		Env:  os.Getenv("ENV"),
		Db:   GetPostgresConfig(),
		Host: os.Getenv("APP_HOST"),
		Port: os.Getenv("APP_PORT"),
		Auth: GetAuthConfig(),
	}
}
