package tui

import (
	"fmt"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/utils"
	"github.com/Dat-one-dev/Sushupti/style"
)

func RenderProjects(daily []data.DailyStat) []string {
	most := utils.MostUsedProject(daily)
	projectTime := utils.ProjectTotalTime(daily, most.ProjectName)
	projectShare := utils.ProjectShareCalc(daily, projectTime)

	content := []string{
		style.Title.Render("Projects"),
		style.Divider,
		"Most Hours  " + style.Highlight.Render(most.ProjectName),
		"Total Time  " + style.Value.Render(utils.FormatTime(projectTime)),
		"Share       " + fmt.Sprintf("%.1f%%", projectShare),
	}

	return content
}
