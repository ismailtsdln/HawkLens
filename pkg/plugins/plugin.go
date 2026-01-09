package plugins

import "context"

// Result represents the data collected by a plugin
type Result struct {
	Platform string                 `json:"platform"`
	DataType string                 `json:"data_type"`
	Data     map[string]interface{} `json:"data"`
}

// Plugin defines the interface that all OSINT modules must implement
type Plugin interface {
	Name() string
	Description() string
	Fetch(ctx context.Context, query string) ([]Result, error)
}
