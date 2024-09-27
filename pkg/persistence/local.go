package persistence

import (
	"fmt"
	"log"
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

	dir := "backup/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}
	}

	file, err := os.Create(fmt.Sprintf("%v/%v-%v.json", dir, database,
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

func (l *Local) Get(database, docs string) error {
	l.logger.Info("getting all docs to restore...")
	return nil
}

func NewLocal() Storage {
	return &Local{
		logger: logger.New(),
	}
}
