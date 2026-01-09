package cli

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fatih/color"
	"github.com/ismailtsdln/HawkLens/internal/analytics"
	"github.com/ismailtsdln/HawkLens/internal/db"
	"github.com/ismailtsdln/HawkLens/internal/engine"
	"github.com/ismailtsdln/HawkLens/pkg/plugins"
	"github.com/olekukonko/tablewriter"
	"github.com/schollz/progressbar/v3"
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
		color.Cyan("\n🔍 Starting multi-platform scan for: %s\n", query)

		pluginNames := plugins.ListPlugins()

		// Initialize Dispatcher with a worker pool size (e.g., 5)
		dispatcher := engine.NewDispatcher(5)
		dispatcher.Run(context.Background())

		// Progress bar setup
		bar := progressbar.NewOptions(len(pluginNames),
			progressbar.OptionSetDescription("Scanning platforms..."),
			progressbar.OptionSetWriter(os.Stderr),
			progressbar.OptionShowCount(),
			progressbar.OptionSetWidth(15),
			progressbar.OptionClearOnFinish(),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "[green]⠿[reset]",
				SaucerHead:    "[green]⠿[reset]",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		)

		startTime := time.Now()

		// Submit jobs to dispatcher
		for _, name := range pluginNames {
			dispatcher.Submit(name, query)
		}

		// Collect results in a separate goroutine
		var allResults []plugins.Result
		var resultsMutex sync.Mutex

		go func() {
			for wrapper := range dispatcher.Results() {
				if wrapper.Error != nil {
					color.Red("[%s] Error: %v\n", wrapper.Platform, wrapper.Error)
					continue
				}
				resultsMutex.Lock()
				allResults = append(allResults, wrapper.Results...)
				resultsMutex.Unlock()
				bar.Add(1)
			}
		}()

		dispatcher.Wait()

		table := tablewriter.NewWriter(os.Stdout)
		table.Append([]string{"Platform", "Type", "Sentiment", "Topics", "Summary"})

		for _, res := range allResults {
			var text string
			switch res.Platform {
			case "twitter":
				text = res.Data["text"].(string)
			case "youtube":
				text = res.Data["title"].(string)
			case "instagram":
				text = res.Data["caption"].(string)
			case "tiktok":
				text = res.Data["hashtag"].(string)
			case "reddit":
				text = res.Data["title"].(string)
			default:
				text = fmt.Sprintf("%v", res.Data)
			}

			analysis := analytics.AnalyzeText(text)

			sentiment := analysis.Sentiment
			if sentiment == "positive" {
				sentiment = color.GreenString(sentiment)
			} else if sentiment == "negative" {
				sentiment = color.RedString(sentiment)
			}

			table.Append([]string{
				color.YellowString(res.Platform),
				res.DataType,
				sentiment,
				fmt.Sprintf("%v", analysis.Topics),
				text,
			})

			if saveToDB {
				pg, err := db.NewPostgresDB("postgres://user:pass@localhost:5432/hawklens?sslmode=disable")
				if err == nil {
					pg.SaveResult(&db.ScanResult{
						Platform: res.Platform,
						DataType: res.DataType,
						Data:     res.Data,
						Query:    query,
					})
					pg.Close()
				}
			}
		}

		table.Render()

		if exportFormat != "" && exportPath != "" {
			var err error
			if exportFormat == "json" {
				err = analytics.ExportToJSON(exportPath, allResults)
			} else if exportFormat == "csv" {
				err = analytics.ExportToCSV(exportPath, allResults)
			}
			if err != nil {
				color.Red("Export error: %v\n", err)
			} else {
				color.Green("\n✅ Data exported to %s\n", exportPath)
			}
		}

		duration := time.Since(startTime)
		color.Cyan("\n🏁 Done! Found %d results in %v across %d platforms.\n", len(allResults), duration, len(pluginNames))
	},
}

func init() {
	scanAllCmd.Flags().StringVarP(&exportFormat, "format", "f", "", "Export format (json|csv)")
	scanAllCmd.Flags().StringVarP(&exportPath, "output", "o", "", "Export output path")
	scanAllCmd.Flags().BoolVarP(&saveToDB, "db", "d", false, "Save results to PostgreSQL database")
	rootCmd.AddCommand(scanAllCmd)
}
