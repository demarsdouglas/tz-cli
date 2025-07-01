package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/demarsdouglas/tz-cli/config"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configured timezones",
	Run: func(cmd *cobra.Command, args []string) {
		zones := config.LoadZones()
		if len(zones) == 0 {
			fmt.Println("No timezones configured.")
			return
		}
		fmt.Println("Configured timezones:")
		for _, zone := range zones {
			fmt.Println(" -", zone)
		}
	},
}
