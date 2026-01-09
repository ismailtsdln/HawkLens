package cli

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ismailtsdln/HawkLens/internal/analytics"
	"github.com/ismailtsdln/HawkLens/pkg/plugins"
	"github.com/spf13/cobra"
)

var (
	exportFormat string
	exportPath   string
	saveToDB     bool
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

		go func() {
			wg.Wait()
			close(resultsChan)
		}()

		var allResults []plugins.Result
		for results := range resultsChan {
			for _, res := range results {
				// Perform NLP analysis
				var text string
				if res.Platform == "twitter" {
					text = res.Data["text"].(string)
				} else if res.Platform == "youtube" {
					text = res.Data["title"].(string)
				} else if res.Platform == "instagram" {
					text = res.Data["caption"].(string)
				}

				analysis := analytics.AnalyzeText(text)
				fmt.Printf("[%s] %s (Sentiment: %s, Topics: %v)\n", res.Platform, res.DataType, analysis.Sentiment, analysis.Topics)

				allResults = append(allResults, res)
			}
		}

		if exportFormat != "" && exportPath != "" {
			var err error
			if exportFormat == "json" {
				err = analytics.ExportToJSON(exportPath, allResults)
			} else if exportFormat == "csv" {
				err = analytics.ExportToCSV(exportPath, allResults)
			}
			if err != nil {
				fmt.Printf("Export error: %v\n", err)
			} else {
				fmt.Printf("\nData exported to %s\n", exportPath)
			}
		}

		duration := time.Since(startTime)
		fmt.Printf("\nDone! Found %d results in %v across %d platforms.\n", len(allResults), duration, len(pluginNames))
	},
}

func init() {
	scanAllCmd.Flags().StringVarP(&exportFormat, "format", "f", "", "Export format (json|csv)")
	scanAllCmd.Flags().StringVarP(&exportPath, "output", "o", "", "Export output path")
	scanAllCmd.Flags().BoolVarP(&saveToDB, "db", "d", false, "Save results to PostgreSQL database")
	rootCmd.AddCommand(scanAllCmd)
}
