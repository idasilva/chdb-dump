package client

import (
	"context"
	"encoding/json"
	"fmt"
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
	rows := c.client.DB(database).AllDocs(context.TODO(), kivik.Param("include_docs", true), kivik.Param("attachments", true))

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
	c.logger.Info("inserting documents into target database...",
		zap.String("dbName", database))

	if err := c.CheckAndCreateDatabase(context.TODO(), database); err != nil {
		return fmt.Errorf("failed to check/create database: %w", err)
	}

	for _, doc := range docs {

		if _, ok := doc.Data["_rev"]; ok {
			delete(doc.Data, "_rev")
		}
		_, _, err := c.client.DB(database).CreateDoc(context.TODO(), doc.Data)
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

func (c *couchdb) CheckAndCreateDatabase(ctx context.Context, dbName string) error {
	// Check if the database exists
	exists, err := c.client.DBExists(ctx, dbName)
	if err != nil {
		c.logger.Warn("Error checking if database exists",
			zap.String("dbName", dbName),
			zap.Error(err))
		return err
	}

	// If the database does not exist, create it
	if !exists {
		if err := c.client.CreateDB(ctx, dbName); err != nil {
			c.logger.Warn("Failed to create database",
				zap.String("dbName", dbName),
				zap.Error(err))
			return err
		}
		c.logger.Info("Database created",
			zap.String("dbName", dbName))
	} else {
		c.logger.Info("Database already exists",
			zap.String("dbName", dbName))
	}

	return nil
}

func New(tp int) Client {
	logger := logger.New()

	var url = os.Getenv("COUCHDB_BACKUP_URL")
	if tp == 2 {
		url = os.Getenv("COUCHDB_RESTORE_URL")
	}
	client, err := kivik.New("couch", url)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}

	return &couchdb{
		client: client,
		logger: logger,
	}
}
