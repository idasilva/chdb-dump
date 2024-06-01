package client

type Document struct {
	ID   string                 `json:"_id"`
	Rev  string                 `json:"_rev"`
	Data map[string]interface{} `json:"data"`
}

type Client interface {
	AllDocs() (string, error)
}
