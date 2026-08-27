package tui

import (
	"fmt"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/charmbracelet/lipgloss"
)

func RenderSidebar(daily []data.DailyStat, symb string) []string {
	total := 0
	today := 0

	for _, day := range daily {
		total += day.TotalSeconds
	}
	if len(daily) > 0 {
		today = daily[len(daily)-1].TotalSeconds
	}

	totalHr := total / 3600
	totalMin := (total % 3600) / 60
	todayHr := today / 3600
	todayMin := (today % 3600) / 60

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4"))

	valueStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6"))

	liveStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("1"))

	content := []string{
		titleStyle.Render("SUSHUPTI"),
		"",
		liveStyle.Render(symb + " LIVE"),
		"",
		titleStyle.Render("TODAY"),
		valueStyle.Render(fmt.Sprintf("%02dh %02dm", todayHr, todayMin)),
		"",
		titleStyle.Render("DATA"),
		fmt.Sprintf("%d days", len(daily)),
		valueStyle.Render(fmt.Sprintf("%02dh %02dm", totalHr, totalMin)),
		"",
		titleStyle.Render("STATUS"),
		liveStyle.Render("● ACTIVE"),
	}
	return content
}
