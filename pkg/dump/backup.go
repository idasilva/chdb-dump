package dump

import (
	"time"

	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type Data struct {
	logger *zap.Logger
}

func (d *Data) Store() {
	defer d.logger.Sync()

	t := time.Now()
	d.logger.Info(t.Format("20060102150405"))

}

func New(opts interface{}) (*Data, error) {

	return &Data{
		logger: logger.New(),
	}, nil
}
