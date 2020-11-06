package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"

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
	Development bool   `env:"DEVELOPMENT" envDefault:"false"`
	Port        int    `env:"PORT" required:"true"`
	TokenSecret string `env:"TOKEN_SECRET" required:"true"`
}

//New returns a pointer to config
func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("No environment variable provided.")
	}
	cfg.PostgresUrl = os.Getenv("POSTGRES_URL")
	fmt.Println("here")
	fmt.Println(cfg.PostgresUrl)
	cfg.ServiceName = AppSrvName


	if cfg.Development {
		fmt.Println("got here!")
		//load a .env file
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatalf("Could not load .env file: %s", err.Error())
		} else {
			log.Info(".env file loaded successfully")
		}

		if err := env.Parse(&cfg); err != nil {
			log.Fatal(err.Error())
		}
	}

	return &cfg, nil
}
