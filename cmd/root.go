package cmd

import (
	"fmt"
	"time"
    "github.com/spf13/cobra"
    "github.com/demarsdouglas/tz-cli/config"
    "github.com/demarsdouglas/tz-cli/timeutil"
)

var live bool

var rootCmd = &cobra.Command{
	Use:   "tz",
	Short: "Display current time in configured timezones",
	Run: func(cmd *cobra.Command, args []string) {
		zones := config.LoadZones()
		if live {
			for {
				fmt.Print("\033[H\033[2J")
				timeutil.PrintZones(zones)
				time.Sleep(1 * time.Second)
			}
		} else {
			timeutil.PrintZones(zones)
		}
	},
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolVarP(&live, "live", "l", false, "Refresh every second")
	rootCmd.AddCommand(addCmd, removeCmd, listCmd, resetCmd)
}
