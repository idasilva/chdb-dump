package client

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"os"

	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	"github.com/idasilva/chdb-dump/pkg/logger"
	"go.uber.org/zap"
)

type couchdb struct {
	client *kivik.Client
	logger *zap.Logger
}

func (c *couchdb) AllDocs(database string) (string, error) {
	c.logger.Info("get all docs...")
	rows := c.client.DB(database).AllDocs(context.TODO(), kivik.IncludeDocs())

	var docs []Document
	for rows.Next() {
		var doc Document
		if err := rows.ScanDoc(&doc.Data); err != nil {
			c.logger.Warn(err.Error())
			continue
		}
		doc.ID, _ = rows.ID()
		doc.Rev, _ = rows.Rev()

		docs = append(docs, doc)
	}

	if err := rows.Err(); err != nil {
		c.logger.Fatal(err.Error())
	}

	jsonData, err := json.MarshalIndent(docs, "", "  ")
	if err != nil {
		c.logger.Fatal(err.Error())
		return "", err
	}

	c.logger.Info("end get all docs...",
		zap.String("data", base64.RawStdEncoding.EncodeToString(jsonData)),
	)

	return string(jsonData), nil

}

func New() Client {
	logger := logger.New()

	client, err := kivik.New("couch", os.Getenv("COUCHDB_URL"))
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	return &couchdb{
		client: client,
		logger: logger,
	}
}
