package cmd

import (
    "github.com/spf13/cobra"
    "github.com/demarsdouglas/tz-cli/config"
    "github.com/demarsdouglas/tz-cli/timeutil"
)

var rootCmd = &cobra.Command{
    Use:   "tz",
    Short: "Display current time in configured timezones",
    Run: func(cmd *cobra.Command, args []string) {
        zones := config.LoadZones()
        timeutil.PrintZones(zones)
    },
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}

func init() {
    rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
}
