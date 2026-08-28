package tui

import (
	"fmt"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/style"
	"github.com/Dat-one-dev/Sushupti/utils"
)

func RenderLeaderboard(daily []data.DailyStat) []string {
	projects := utils.ProjectLeaderboard(daily)

	content := []string{
		style.Title.Render("Projects Leaderboard"),
		style.Divider,
	}

	for i, project := range projects {
		line := fmt.Sprintf(
			"%d. %-20s %s",
			i+1,
			project.ProjectName,
			utils.FormatTime(project.TotalSeconds),
		)

		content = append(content,
			style.ProjectColors[i%len(style.ProjectColors)].Render(line),
		)
	}

	return content
}
