package deps

import (
	"github.com/isongjosiah/lernen-api/config"
	"github.com/isongjosiah/lernen-api/dal"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Dependencies struct
type Dependencies struct {
	// DAL
	DAL *dal.DAL
}

// New Dependencies instance
func New(cfg *config.Config) (*Dependencies, error) {
	dal, err := dal.New(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to setup DAL")
	}
	log.Info("DAL: ok")

	deps := &Dependencies{
		DAL: dal,
	}

	return deps, nil
}
