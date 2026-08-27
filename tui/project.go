package tui

import (
	"fmt"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/utils"
	"github.com/charmbracelet/lipgloss"
)

func RenderProjects(daily []data.DailyStat) []string {
	most := utils.MostUsedProject(daily)
	projectTime := utils.ProjectTotalTime(daily, most.ProjectName)
	projectShare := utils.ProjectShareCalc(daily, projectTime)

	projectsTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4")).
		Render("Projects")

	mostProjectName := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("1")).
		Render(most.ProjectName)

	shareTxt := fmt.Sprintf("%.1f%%", projectShare)

	content := []string{
		projectsTitle,
		"──────────────────────────────────────────",
		"Most Hours  " + mostProjectName,
		"Total Time  " + fmt.Sprintf("%dh %02dm",
			projectTime/3600,
			(projectTime%3600)/60),
		"Share       " + shareTxt,
	}

	return content
}
