package config

import (
	"os"

	"github.com/adityatresnobudi/restapi/pkg/constants"
)

type Config struct {
	Postgres PostgresConfig
	Http     HttpConfig
}

type PostgresConfig struct {
	Port     string
	Host     string
	User     string
	Password string
	DBName   string
}

type HttpConfig struct {
	Port string
	Host string
}

func NewConfig() Config {
	cfg := Config{
		Http: HttpConfig{
			Port: os.Getenv(constants.HTTPPort),
			Host: os.Getenv(constants.APIHost),
		},
		Postgres: PostgresConfig{
			Port: os.Getenv(constants.DBPort),
			Host: os.Getenv(constants.DBHost),
			User: os.Getenv(constants.DBUser),
			Password: os.Getenv(constants.DBPassword),
			DBName: os.Getenv(constants.DBName),
		},
	}

	return cfg
}