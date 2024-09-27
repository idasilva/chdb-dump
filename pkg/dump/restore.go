package dump

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/idasilva/chdb-dump/pkg/client"
	"github.com/idasilva/chdb-dump/pkg/context"
	"github.com/idasilva/chdb-dump/pkg/persistence"

	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type DataTo struct {
	logger  *zap.Logger
	client  client.Client
	storage persistence.Storage
}

func (d *DataTo) Exec() error {
	d.logger.Info("start restore data...")

	docs := make(map[string][]client.Document)

	err := filepath.Walk("backup", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			d.logger.Warn("Failed to read file",
				zap.String("filePath", path),
				zap.Error(err))
			return nil
		}

		var docArray []client.Document
		if err := json.Unmarshal(content, &docArray); err != nil {
			d.logger.Warn("Failed to unmarshal JSON",
				zap.String("filePath", path),
				zap.Error(err))
			return nil
		}

		baseName := strings.Split(info.Name(), "-")[0]
		docs[baseName] = append(docs[baseName], docArray...)

		d.logger.Info("Documents loaded from file",
			zap.String("baseName", baseName),
			zap.Int("docCount", len(docArray)))
		return nil
	})

	if err != nil {
		return nil
	}

	d.logger.Info("len of documents",
		zap.Int("docCount", len(docs)))

	// _ = d.client.PutDocs("", docs)
	// if err != nil {
	// 	return err
	// }
	return nil

}

func NewRestore(context *context.Context) (*DataTo, error) {
	logger := logger.New()

	logger.Info("starting configuring restore..",
		zap.String("type", context.Persistence),
	)

	loadFunction := func(persistence persistence.Storage) (*DataTo, error) {
		return &DataTo{
			logger:  logger,
			storage: persistence,
			client:  client.New(),
		}, nil
	}

	if context.Persistence == persistence.LC {
		return loadFunction(persistence.NewLocal())
	}

	return loadFunction(persistence.NewS3())
}
