package cli

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ismailtsdln/HawkLens/pkg/plugins"
	"github.com/spf13/cobra"
)

var scanAllCmd = &cobra.Command{
	Use:   "scan [query]",
	Short: "Scan all platforms concurrently for a specific query",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		fmt.Printf("Starting multi-platform scan for: %s\n\n", query)

		pluginNames := plugins.ListPlugins()
		var wg sync.WaitGroup
		resultsChan := make(chan []plugins.Result, len(pluginNames))

		startTime := time.Now()

		for _, name := range pluginNames {
			wg.Add(1)
			go func(pluginName string) {
				defer wg.Done()
				p, _ := plugins.GetPlugin(pluginName)
				res, err := p.Fetch(context.Background(), query)
				if err != nil {
					fmt.Printf("[%s] Error: %v\n", pluginName, err)
					return
				}
				resultsChan <- res
			}(name)
		}

		// Close channel when all goroutines finish
		go func() {
			wg.Wait()
			close(resultsChan)
		}()

		totalResults := 0
		for results := range resultsChan {
			for _, res := range results {
				fmt.Printf("[%s] %s found data: %v\n", res.Platform, res.DataType, res.Data)
				totalResults++
			}
		}

		duration := time.Since(startTime)
		fmt.Printf("\nDone! Found %d results in %v across %d platforms.\n", totalResults, duration, len(pluginNames))
	},
}

func init() {
	rootCmd.AddCommand(scanAllCmd)
}
