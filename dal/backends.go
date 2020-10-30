package dal

import (
	"github.com/isongjosiah/lernen-api/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//Backends struct contains all the services utilized at the backend of the api.
type Backends struct {
	Postgres *sqlx.DB
}

func NewBackends(cfg *config.Config) (*Backends, error) {
	b := &Backends{}

	if err := b.connectPostgres(cfg); err != nil {
		return nil, err
	}
	return b, nil
}

func (b Backends) connectPostgres(cfg *config.Config) error {
	//TODO(josiah): find out how to setup postgres with docker, Using ElephantSQl just to test the code now.
	log.Info("Postgres Client: connecting to host ...")
	db, err := sqlx.Connect("postgres", cfg.PostgresUrl)
	if err != nil {
		err := errors.Wrapf(err, "Postgres: unable to open an initial connection to client: %v", cfg.PostgresUrl)
		return err
	}

	b.Postgres = db
	return nil
}
