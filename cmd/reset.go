package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/demarsdouglas/tz-cli/config"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Delete all configured timezones",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Remove(config.ConfigFilePath())
		if err != nil {
			fmt.Println("Error deleting config:", err)
		} else {
			fmt.Println("All timezones reset.")
		}
	},
}
