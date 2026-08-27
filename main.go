package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func logError(err error) bool {
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}

func main() {
	startDate := flag.String("s", "", "start date")
	endDate := flag.String("e", "", "end date")
	darkBool := flag.Bool("dark", false, "dark mode")
	flag.Parse()
	if *startDate == "" && *endDate == "" {
		today := time.Now()
		*endDate = today.Format("02-01-06")
		*startDate = today.AddDate(0, 0, -10).Format("02-01-06")
	}
	parsed, err := utils.ParseTime(*startDate)
	if !logError(err) {
		return
	}
	parsedE, err := utils.ParseTime(*endDate)
	if !logError(err) {
		return
	}

	config := data.Config{}
	user, err := user.Current()
	if !logError(err) {
		return
	}

	configPath := filepath.Join(user.HomeDir, ".wakatime.cfg")
	configData, err := os.ReadFile(configPath)
	if !logError(err) {
		return
	}

	lines := strings.Split(string(configData), "\n")

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "api_url" {
			config.APIUrl = fmt.Sprintf("%s/users/current/summaries?start=%s&end=%s", value, parsed, parsedE)
		}

		if key == "api_key" {
			config.APIkey = value
		}
	}
	var dailyStats []data.DailyStat

	httpResponse, err := utils.SendRequest(config)
	if !logError(err) {
		return
	} else {
		defer httpResponse.Body.Close()
		x, err := io.ReadAll(httpResponse.Body)
		fmt.Println(string(x))
		if !logError(err) {
			return
		}
		response := data.Response{}
		err = json.Unmarshal(x, &response)
		if !logError(err) {
			return
		}
		fmt.Println("LANGUAGES:")

		for _, day := range response.Data {
			fmt.Println(day.Range.Date, day.Languages)
		}
		for _, day := range response.Data {
			dailyStats = append(dailyStats, data.DailyStat{
				Date:         day.Range.Date,
				TotalSeconds: day.GrandTotal.TotalSeconds,
				Projects:     day.Projects,
				Editors:      day.Editors,
				Languages:    day.Languages,
			})
		}

		utils.Graph(dailyStats)
		err = utils.ExportGraph(dailyStats, "sushupti.png", *darkBool)
		if !logError(err) {
			return
		}

	}

	p := tea.NewProgram(model{
		daily: dailyStats,
	})
	if _, err := p.Run(); !logError(err) {
		return
	}
}
