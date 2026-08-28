package tui

import (
	"fmt"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/style"
	"github.com/Dat-one-dev/Sushupti/utils"
)

func RenderSidebar(daily []data.DailyStat, symb string) []string {
	total := utils.TotalTime(daily)
	today := 0

	if len(daily) > 0 {
		today = daily[len(daily)-1].TotalSeconds
	}

	content := []string{
		style.Title.Render("SUSHUPTI"),
		"",
		style.Highlight.Render(symb + " LIVE"),
		"",
		style.Title.Render("TODAY"),
		style.Value.Render(utils.FormatTime(today)),
		"",
		style.Title.Render("DATA"),
		fmt.Sprintf("%d days", len(daily)),
		style.Value.Render(utils.FormatTime(total)),
		"",
		style.Title.Render("STATUS"),
		style.Highlight.Render("● ACTIVE"),
	}

	return content
}
