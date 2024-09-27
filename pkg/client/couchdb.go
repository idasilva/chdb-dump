package client

import (
	"context"
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
	c.logger.Info("getting all docs...",
		zap.String("dbName", database))
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

	c.logger.Info("end get all docs...")

	return string(jsonData), nil
}

func (c *couchdb) PutDocs(database string, docs []Document) error {
	c.logger.Info("Inserting documents into target database...",
		zap.String("dbName", database))

	for _, doc := range docs {
		// Insert each document into the target database
		_, err := c.client.DB(database).Put(context.TODO(), doc.ID, doc.Data)
		if err != nil {
			c.logger.Warn("Failed to insert document",
				zap.String("docID", doc.ID),
				zap.String("targetDatabase", database),
				zap.Error(err))
			continue
		}

		c.logger.Info("Document inserted successfully",
			zap.String("docID", doc.ID),
			zap.String("targetDatabase", database))
	}

	c.logger.Info("All documents inserted successfully.")
	return nil
}

func New() Client {
	logger := logger.New()

	client, err := kivik.New("couch", os.Getenv("COUCHDB_BACKUP_URL"))
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	return &couchdb{
		client: client,
		logger: logger,
	}
}
