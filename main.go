package main

import (
	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type refreshMsg []data.DailyStat

func main() {
	start, end, err := utils.DateRange()
	if !utils.LogError(err) {
		return
	}

	config, err := utils.LoadConfig(start, end)
	if !utils.LogError(err) {
		return
	}

	daily, err := utils.FetchDaily(config)
	if !utils.LogError(err) {
		return
	}

	p := tea.NewProgram(model{
		daily: daily,
	})

	_, err = p.Run()
	utils.LogError(err)
}
