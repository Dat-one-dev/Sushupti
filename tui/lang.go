package tui

import (
	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/style"
	"github.com/Dat-one-dev/Sushupti/utils"
)

func ProjectBar(daily []data.DailyStat, width, anim int) []string {
	projects := utils.ProjectLeaderboard(daily)
	total := utils.TotalTime(daily)
	bar := ""
	layout := make([]int, 0, width)

	content := []string{
		style.Title.Render("Projects Activity") + "  " +
			style.Value.Render(utils.FormatTime(total)),
		style.Divider,
	}

	if len(projects) == 0 || total == 0 {
		content = append(content, "─")
		return content
	}

	for i, project := range projects {
		size := project.TotalSeconds * width / total

		for j := 0; j < size; j++ {
			if len(layout) >= width {
				break
			}

			layout = append(layout, i%len(style.ProjectColors))
		}
	}

	if anim > len(layout) {
		anim = len(layout)
	}

	if anim < 0 {
		anim = 0
	}

	for i := 0; i < anim; i++ {
		bar += style.ProjectColors[layout[i]].Render("#")
	}
	content = append(content, bar)
	return content
}
