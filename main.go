package main

import (
	"github.com/isongjosiah/lernen-api/api"
	"github.com/isongjosiah/lernen-api/config"
	"github.com/isongjosiah/lernen-api/deps"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	allowConnectionsAfterShutdown = 5 * time.Second
)

func main() {
	// application config. for now we are using the .env file in the config directory
	cfg, err := config.New()
	if err != nil {
		log.Warn("Could not set up the environment")
	} else {
		log.Info("ENV: ok")
	}

	// set up the project dependencies
	deps, err := deps.New(cfg)
	if err != nil {
		log.Fatalf("Unable to setup dependencies, Error: %s", err.Error())
	}
	log.Info("Deps: ok")

	//	start the api server
	a := &api.API{
		Config: cfg,
		Deps:   deps,
	}

	log.Fatal(a.Serve())
}
