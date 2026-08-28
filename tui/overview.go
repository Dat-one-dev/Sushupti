package tui

import (
	"fmt"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/style"
	"github.com/Dat-one-dev/Sushupti/utils"
)

func RenderOverview(daily []data.DailyStat, anim int) []string {
	//ALL VARIABLES ARE HERE
	activeDays := utils.ActiveDays(daily)
	bestDay := utils.BestDay(daily)
	dailyAverage := utils.DailyAvg(daily)
	total := utils.TotalTime(daily)
	maxSec := bestDay.TotalSeconds

	//contetblcok
	content := []string{
		style.Title.Render("Overview - Last 10 Days"),
		style.Divider,
		"Total       " + style.Value.Render(utils.FormatTime(total)),
		"Active Days " + style.Value.Render(fmt.Sprintf("%d", activeDays)),
	}

	if bestDay.TotalSeconds > 0 {
		content = append(content, "Best Day    "+style.Value.Render(bestDay.Date[5:])+"  "+style.Value.Render(utils.FormatTime(bestDay.TotalSeconds)))
	} else {
		content = append(content, "Best Day    "+style.Value.Render("No data"))
	}

	content = append(content, "Daily Avg   "+style.Value.Render(utils.FormatTime(dailyAverage)), "")

	for _, day := range daily {
		line := day.Date[5:] + "  "
		line += style.Value.Render(utils.ProgBar(
			day.TotalSeconds,
			maxSec,
			20,
			anim,
		))

		line += "  " + utils.FormatTime(day.TotalSeconds)
		content = append(content, line)
	}

	return content
}
