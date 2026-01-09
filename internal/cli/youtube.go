package cli

import (
	"context"
	"fmt"

	"github.com/ismailtsdln/HawkLens/internal/plugins/youtube"
	"github.com/spf13/cobra"
)

var youtubeCmd = &cobra.Command{
	Use:   "youtube [query]",
	Short: "Scan YouTube for a specific query",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		fmt.Printf("Starting YouTube scan for: %s\n", query)

		plugin := youtube.NewYouTubePlugin("mock_api_key")
		results, err := plugin.Fetch(context.Background(), query)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		for _, res := range results {
			fmt.Printf("[%s] %s: %v\n", res.Platform, res.DataType, res.Data["title"])
		}
	},
}

func init() {
	rootCmd.AddCommand(youtubeCmd)
}
