package client

import (
	"context"

	"github.com/go-kivik/kivik/v4"
)

type couchdb struct {
	client *kivik.Client
}

func (c *couchdb) AllDocs() {

	_ = c.client.DB("").AllDocs(context.TODO())

}

func New() Client {

	client, err := kivik.New("couch", "http://localhost:5984/")
	if err != nil {
		panic(err)
	}

	return &couchdb{
		client: client,
	}
}
