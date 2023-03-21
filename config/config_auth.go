package config

import (
	"os"
)

type AuthConfiguration struct {
	PublicPemPath string `env:"AUTH_PUBLIC_PEM_PATH"`
}

func GetAuthConfig() AuthConfiguration {

	return AuthConfiguration{
		PublicPemPath: os.Getenv("AUTH_PUBLIC_PEM_PATH"),
	}
}
