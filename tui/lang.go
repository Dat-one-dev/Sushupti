package tui

import (
	"fmt"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/utils"
	"github.com/charmbracelet/lipgloss"
)

func RenderLang(daily []data.DailyStat) []string {
	projects := utils.ProjectLeaderboard(daily)
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("4")).Render("Projects Activity")
	barStyl := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("6"))

	content := []string{
		title,
		"──────────────────────────────────────────",
	}

	if len(projects) == 0 {
		content = append(content, "No Projects Data")
		return content
	}
	max := projects[0].TotalSeconds

	for _, project := range projects {
		bar := 0

		if max > 0 {
			bar = project.TotalSeconds * 20 / max
		}
		line := ""

		for i := 0; i < bar; i++ {
			line += barStyl.Render("█")
		}

		hrs := project.TotalSeconds / 3600
		mns := (project.TotalSeconds % 3600) / 60

		line += fmt.Sprintf(
			" %-15s %dh %02dm", project.ProjectName, hrs, mns,
		)

		content = append(content, line)
	}
	return content
}
