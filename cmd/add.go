package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/demarsdouglas/tz-cli/config"
)

var addCmd = &cobra.Command{
    Use:   "add [timezone]",
    Short: "Add a timezone",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        zone := args[0]
        if err := config.AddZone(zone); err != nil {
            fmt.Println("Failed to add timezone:", err)
        } else {
            fmt.Println("Added timezone:", zone)
        }
    },
}
