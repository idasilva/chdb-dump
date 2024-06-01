package dump

import (
	"time"

	"github.com/idasilva/chdb-dump/pkg/context"
	"github.com/idasilva/chdb-dump/pkg/persistence"

	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type Data struct {
	logger  *zap.Logger
	storage persistence.Storage
}

func (d *Data) Backup() {
	defer d.logger.Sync()

	t := time.Now()
	d.logger.Info(t.Format("20060102150405"))

	//# Grab our data from couchdb

	// # Check for export errors

}

func (d *Data) Store() {
	defer d.logger.Sync()

	t := time.Now()
	d.logger.Info(t.Format("20060102150405"))

}

func New(context *context.Context) (*Data, error) {
	logger := logger.New()

	logger.Info("starting configuring backup using persistence...",
		zap.String("type", context.Persistence),
	)

	if context.Persistence == persistence.LC {
		return &Data{
			logger:  logger,
			storage: persistence.NewLocal(),
		}, nil
	}

	return &Data{
		logger:  logger,
		storage: persistence.NewS3(),
	}, nil
}
