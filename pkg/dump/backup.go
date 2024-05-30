package dump

import (
	"time"

	"github.com/idasilva/chdb-dump/pkg/context"

	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type Data struct {
	logger *zap.Logger
}

func (d *Data) Backup() {
	defer d.logger.Sync()

	t := time.Now()
	d.logger.Info(t.Format("20060102150405"))

}

func (d *Data) Store() {
	defer d.logger.Sync()

	t := time.Now()
	d.logger.Info(t.Format("20060102150405"))

}

func New(context *context.Context) (*Data, error) {
	logger := logger.New()

	logger.Info(context.Persistence)

	return &Data{
		logger: logger,
	}, nil
}
