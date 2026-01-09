package instagram

import (
	"context"
	"fmt"

	"github.com/ismailtsdln/HawkLens/pkg/plugins"
)

type InstagramPlugin struct{}

func NewInstagramPlugin() *InstagramPlugin {
	return &InstagramPlugin{}
}

func (i *InstagramPlugin) Name() string {
	return "instagram"
}

func (i *InstagramPlugin) Description() string {
	return "Scrapes public Instagram profiles and posts."
}

func (i *InstagramPlugin) Fetch(ctx context.Context, query string) ([]plugins.Result, error) {
	fmt.Printf("[Instagram] Fetching data for query: %s\n", query)

	results := []plugins.Result{
		{
			Platform: "instagram",
			DataType: "post",
			Data: map[string]interface{}{
				"id":       "insta_123",
				"username": "hawk_lens_osint",
				"caption":  "Visualizing social data #osint",
				"likes":    150,
			},
		},
	}

	return results, nil
}
