package tui

import (
	"fmt"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/charmbracelet/lipgloss"
)

func RenderOverview(daily []data.DailyStat, anim int) []string {
	//ALL VARIABLES ARE HERE
	maxSec := 0
	total := 0
	bestDay := ""
	bestTime := 0
	activeDays := 0
	dailyAverage := 0

	//ALL LIPGLLOSS STUFF IS HERE
	overviewTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4")).
		Render("Overview - Last 10 Days")

	totalStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6"))

	charBar := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6"))

	//doing ts again for 6-7th time instead of making a better way to do it

	// max seconds loop
	for _, day := range daily {
		if day.TotalSeconds > maxSec {
			maxSec = day.TotalSeconds
		}
	}
	//total seconds loop
	for _, day := range daily {
		total += day.TotalSeconds
	}

	//activeDays loop
	for _, day := range daily {
		if day.TotalSeconds > 0 {
			activeDays++
		}
	}

	//bestDay bestTime loop
	for _, day := range daily {
		if day.TotalSeconds > bestTime {
			bestTime = day.TotalSeconds
			bestDay = day.Date
		}
	}

	//Dailyavg
	if activeDays > 0 {
		dailyAverage = total / activeDays
	}

	content := []string{
		overviewTitle,
		"──────────────────────────────────────────",
		"Total       " + totalStyle.Render(fmt.Sprintf("%dh %02dm", total/3600, (total%3600)/60)),
		"Active Days " + totalStyle.Render(fmt.Sprintf("%d", activeDays)),
		"Best Day    " + totalStyle.Render(bestDay[5:]) + "  " +
			totalStyle.Render(fmt.Sprintf("%dh %02dm", bestTime/3600, (bestTime%3600)/60)),
		"Daily Avg   " + totalStyle.Render(fmt.Sprintf("%dh %02dm", dailyAverage/3600, (dailyAverage%3600)/60)),
		"",
	}

	for _, day := range daily {
		bar := 0

		if maxSec > 0 {
			bar = day.TotalSeconds * 20 / maxSec

			if bar > anim {
				bar = anim
			}
		}

		line := day.Date[5:] + "  "

		if bar == 0 {
			line += "─"
		} else {
			for j := 0; j < bar; j++ {
				line += charBar.Render("#")
			}
		}

		hours := day.TotalSeconds / 3600
		minutes := (day.TotalSeconds % 3600) / 60

		line += fmt.Sprintf("  %dh %02dm", hours, minutes)

		content = append(content, line)
	}

	return content
}
