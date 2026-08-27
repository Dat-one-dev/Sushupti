package tui

import (
	"fmt"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/utils"
	"github.com/charmbracelet/lipgloss"
)

func RenderLeaderboard(daily []data.DailyStat) []string {
	projects := utils.ProjectLeaderboard(daily)

	leaderboardTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4")).
		Render("Projects Leaderboard")

	content := []string{
		leaderboardTitle,
		"──────────────────────────────────────────",
	}

	for i, project := range projects {
		hours := project.TotalSeconds / 3600
		minutes := (project.TotalSeconds % 3600) / 60

		content = append(content, fmt.Sprintf(
			"%d. %-20s %dh %02dm",
			i+1,
			project.ProjectName,
			hours,
			minutes,
		))
	}

	return content
}
