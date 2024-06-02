package dump

import (
	"github.com/idasilva/chdb-dump/pkg/client"
	"github.com/idasilva/chdb-dump/pkg/context"
	"github.com/idasilva/chdb-dump/pkg/persistence"

	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type Data struct {
	logger   *zap.Logger
	client   client.Client
	storage  persistence.Storage
	database string
}

func (d *Data) Exec() error {
	d.logger.Info("start backup data...")

	docs, err := d.client.AllDocs(d.database)
	if err != nil {
		return err
	}

	err = d.storage.Store(d.database, docs)
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

	if context.Database == "" {
		return nil, errDatabaseNotFound
	}

	loadFunction := func(persistence persistence.Storage) (*Data, error) {
		return &Data{
			logger:   logger,
			storage:  persistence,
			client:   client.New(),
			database: context.Database,
		}, nil
	}

	if context.Persistence == persistence.LC {
		return loadFunction(persistence.NewLocal())
	}

	return loadFunction(persistence.NewS3())
}
