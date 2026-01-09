package cli

import (
	"context"
	"fmt"

	"github.com/ismailtsdln/HawkLens/internal/plugins/twitter"
	"github.com/spf13/cobra"
)

var twitterCmd = &cobra.Command{
	Use:   "twitter [query]",
	Short: "Scan Twitter for a specific query",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		fmt.Printf("Starting Twitter scan for: %s\n", query)

		plugin := twitter.NewTwitterPlugin("mock_api_key")
		results, err := plugin.Fetch(context.Background(), query)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		for _, res := range results {
			fmt.Printf("[%s] %s: %v\n", res.Platform, res.DataType, res.Data["text"])
		}
	},
}

func init() {
	rootCmd.AddCommand(twitterCmd)
}
