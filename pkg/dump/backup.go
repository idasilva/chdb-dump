package dump

import (
	"strings"

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
	database []string
}

func (d *Data) Exec() error {
	d.logger.Info("start backup data...",
		zap.String("databases", strings.Join(d.database, "")))
	for _, db := range d.database {
		docs, err := d.client.AllDocs(db)
		if err != nil {
			return err
		}

		err = d.storage.Store(db, docs)
		if err != nil {
			return err
		}

		d.logger.Info("end backup data...",
			zap.String("dbName", db),
		)
	}
	return nil

}

func NewBackup(context *context.Context) (*Data, error) {
	logger := logger.New()

	logger.Info("starting configuring backup using persistence...",
		zap.String("type", context.Persistence),
	)

	if len(context.Database) == 0 {
		return nil, errDatabaseNotFound
	}

	loadFunction := func(persistence persistence.Storage) (*Data, error) {
		return &Data{
			logger:   logger,
			storage:  persistence,
			client:   client.New(1),
			database: context.Database,
		}, nil
	}

	if context.Persistence == persistence.LC {
		return loadFunction(persistence.NewLocal())
	}

	return loadFunction(persistence.NewS3())
}
