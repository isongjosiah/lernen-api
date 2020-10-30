package dal

import (
	"github.com/isongjosiah/lernen-api/config"
	"github.com/pkg/errors"
)

//DAL is and object housing the entire data access layer
type DAL struct {
	Backends *Backends

	// DAL Objects
	UserDAL IUserDAL
}

func (d *DAL) setupBackends(cfg *config.Config) error {
	backends, err := NewBackends(cfg)
	if err != nil {
		return errors.Wrapf(err, "Backends set up failed")
	}
	d.Backends = backends

	return nil
}

func (d *DAL) setupDALObjects(cfg *config.Config) error {

	return nil
}

// New creates, configures and returns a new DAL object
func New(cfg *config.Config) (*DAL, error) {
	dal := &DAL{}

	if err := dal.setupBackends(cfg); err != nil {
		return nil, err
	}

	if err := dal.setupDALObjects(cfg); err != nil {
		return nil, err
	}

	return dal, nil
}
