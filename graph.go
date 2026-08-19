package main

import "fmt"

func graph(dailyStats []DailyStat) {
	maxSec := 0
	for _, stat := range dailyStats {
		if stat.TotalSeconds > maxSec {
			maxSec = stat.TotalSeconds
		}
	}

	for _, stat := range dailyStats {
		barLength := 0
		if maxSec > 0 {
			barLength = stat.TotalSeconds * 30 / maxSec
		}
		bar := ""
		for i := 0; i < barLength; i++ {
			bar += "█"
		}
		hours := stat.TotalSeconds / 3600
		min := (stat.TotalSeconds % 3600) / 60

		fmt.Printf("%s | %-30s %dh %02dm\n", stat.Date, bar, hours, min)
	}
}
