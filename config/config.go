package config

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/caarlos0/env.v2"
)

const (
	AppSrvName = "lernen-api"
)

//Config contains the necessary configuration for the file
type Config struct {
	ServiceName string
	PostgresUrl string `env:"POSTGRES_URL" required:"true"`
	Development bool   `env:"DEVELOPMENT" envDefault:"true"`
}

//New returns a pointer to config
func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err.Error())
	}

	cfg.ServiceName = AppSrvName

	if cfg.Development {
		//load a .env file
		err := godotenv.Load("./.env")
		if err != nil {
			err := errors.Wrap(err, "Error loading the environment: .env")
			return &cfg, err //you still have to return config otherwise you get a nil pointer dereference error
		} else {
			log.Info(".env file loaded successfully")
		}

		if err := env.Parse(&cfg); err != nil {
			log.Fatal(err.Error())
		}
	}
	return &cfg, nil
}
