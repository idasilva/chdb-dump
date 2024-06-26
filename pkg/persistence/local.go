package persistence

import (
	"fmt"
	"os"
	"time"

	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type Local struct {
	logger *zap.Logger
}

func (l *Local) Store(database, docs string) error {
	l.logger.Info("writing docs to output file...")

	file, err := os.Create(fmt.Sprintf("%v-%v.json", database,
		time.Now().Format("2006-02-01")))
	if err != nil {
		l.logger.Fatal("FATAL")
		return err
	}

	defer file.Close()

	_, err = file.Write([]byte(docs))
	if err != nil {
		l.logger.Fatal("FATAL")
		return err
	}

	l.logger.Info("docs successfully written to output file...")

	return nil
}

func NewLocal() Storage {
	return &Local{
		logger: logger.New(),
	}
}
