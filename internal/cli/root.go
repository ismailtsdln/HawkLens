package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hawklens",
	Short: "HawkLens is a multi-platform social OSINT & analytics framework",
	Long: `HawkLens is a modern, high-performance OSINT tool designed to collect and analyze 
data from various social media platforms like Twitter, YouTube, Reddit, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to HawkLens! Use --help to see available commands.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be defined here
}
