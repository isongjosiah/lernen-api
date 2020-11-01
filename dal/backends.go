package dal

import (
	"github.com/isongjosiah/lernen-api/config"
	"github.com/isongjosiah/lernen-api/dal/model"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres imported for side effects
)

//Backends struct contains all the services utilized at the backend of the api.
type Backends struct {
	Database *gorm.DB
}

func NewBackends(cfg *config.Config) (*Backends, error) {
	b := &Backends{}

	if err := b.connectDatabase(cfg); err != nil {
		return nil, err
	}
	return b, nil
}

func (b Backends) connectDatabase(cfg *config.Config) error {
	log.Info("Database : connecting to client ...")
	db, err := gorm.Open("postgres", cfg.PostgresUrl)
	if err != nil {
		err := errors.Wrapf(err, "Database: unable to open an initial connection to ElephantSQL client: %v", cfg.PostgresUrl)
		return err
	}

	b.Database = db
	b.Database.Debug().AutoMigrate(&model.User{}) // database migration. No idea how this would be useful. TODO(josiah): research on gorm's database migration and it's usefulness for the project
	return nil
}
