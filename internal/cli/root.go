package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var banner = `
  _    _                _   _                      
 | |  | |              | | | |                     
 | |__| | __ _ __      | | | | ___ _ __  ___       
 |  __  |/ _' | \ \ /\ / / | |/ _ \ '_ \/ __|      
 | |  | | (_| |  \ V  V /  | |  __/ | | \__ \      
 |_|  |_|\__,_|   \_/\_/   |_|\___|_| |_|___/      
                                                   
   Multi-Platform Social OSINT & Analytics
`

var rootCmd = &cobra.Command{
	Use:   "hawklens",
	Short: "HawkLens is a multi-platform social OSINT & analytics framework",
	Long:  color.CyanString(banner) + "\nHawkLens is a modern, high-performance OSINT tool designed to collect and analyze data from various social media platforms.",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan(banner)
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
