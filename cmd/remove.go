package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/demarsdouglas/tz-cli/config"
)

var removeCmd = &cobra.Command{
	Use:   "remove [timezone]",
	Short: "Remove a timezone from the config",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		zone := args[0]
		if err := config.RemoveZone(zone); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Removed:", zone)
		}
	},
}
