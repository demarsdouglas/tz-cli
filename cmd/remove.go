package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"github.com/demarsdouglas/tz-cli/config"
	"github.com/demarsdouglas/tz-cli/timeutil"
)

var removeCmd = &cobra.Command{
	Use:   "remove [timezone]",
	Short: "Remove a timezone",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := strings.Join(args, " ")

		zone, err := timeutil.FindMatchingZone(input)
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		if err := config.RemoveZone(zone); err != nil {
			fmt.Println("❌ Failed to remove timezone:", err)
		} else {
			fmt.Println("✅ Removed timezone:", zone)
		}
	},
}
