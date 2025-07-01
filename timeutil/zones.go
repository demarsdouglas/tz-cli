package timeutil

import (
	"fmt"
	"time"
)

func PrintZones(zones []string) {
	now := time.Now()

	fmt.Println("🕒 Timezone Overview:\n---------------------------")

	fmt.Printf("%-20s %s\n", "Local", now.Format("2006-01-02 15:04:05"))

	for _, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			fmt.Printf("%-20s ❌ Invalid timezone\n", zone)
			continue
		}
		tzTime := now.In(loc)
		fmt.Printf("%-20s %s\n", zone, tzTime.Format("2006-01-02 15:04:05"))
	}
}
