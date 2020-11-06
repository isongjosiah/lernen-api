package dal

import (
	"github.com/isongjosiah/lernen-api/config"
	"github.com/isongjosiah/lernen-api/dal/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
)

//DAL is and object housing the entire data access layer
type DAL struct {
	// DAL Objects
	UserDAL IUserDAL
}

//TODO(josiah): definitely refactor this code.
func (d *DAL) setupDALObjects(cfg *config.Config) error {
	log.Infof("JwT secret is %s", os.Getenv("TOKEN_SECRET"))
	log.Info("Database : connecting to client ...")
	db, err := gorm.Open("postgres", cfg.PostgresUrl)
	if err != nil {
		err := errors.Wrapf(err, "Database: unable to open an initial connection to ElephantSQL client: %v", cfg.PostgresUrl)
		return err
	}

	db.Debug().AutoMigrate(&model.User{}) // database migration. No idea how this would be useful. TODO(josiah): research on gorm's database migration and it's usefulness for the project

	u := NewUserDAL()
	u.SetUp(db)
	d.UserDAL = u
	return nil
}

// New creates, configures and returns a new DAL object
func New(cfg *config.Config) (*DAL, error) {
	dal := &DAL{}

	if err := dal.setupDALObjects(cfg); err != nil {
		return nil, err
	}

	return dal, nil
}
