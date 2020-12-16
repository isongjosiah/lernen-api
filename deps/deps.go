package deps

import (
	"github.com/isongjosiah/lernen-api/config"
	"github.com/isongjosiah/lernen-api/dal"
	"github.com/isongjosiah/lernen-api/services"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Dependencies struct
type Dependencies struct {
	//Services
	AWS *services.AWS

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

	aws, err := services.NewAWS(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to setup AWS")
	}
	log.Info("AWS:ok")

	deps := &Dependencies{
		AWS: aws,
		DAL: dal,
	}

	return deps, nil
}
