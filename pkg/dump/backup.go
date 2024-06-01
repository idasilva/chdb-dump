package dump

import (
	"github.com/idasilva/chdb-dump/pkg/client"
	"github.com/idasilva/chdb-dump/pkg/context"
	"github.com/idasilva/chdb-dump/pkg/persistence"

	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type Data struct {
	logger  *zap.Logger
	client  client.Client
	storage persistence.Storage
}

func (d *Data) Exec() error {
	d.logger.Info("start backup data...")

	docs, err := d.client.AllDocs()
	if err != nil {
		return err
	}

	err = d.storage.Store(docs)
	if err != nil {
		return err
	}

	d.logger.Info("end backup data...")

	return nil

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
			client:  client.New(),
		}, nil
	}

	return &Data{
		logger:  logger,
		storage: persistence.NewS3(),
		client:  client.New(),
	}, nil
}
