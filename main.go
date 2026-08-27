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

	tea "github.com/charmbracelet/bubbletea"
)

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
	parsed, err := parseTime(*startDate)
	if !logError(err) {
		return
	}
	parsedE, err := parseTime(*endDate)
	if !logError(err) {
		return
	}

	config := Config{}
	user, err := user.Current()
	if !logError(err) {
		return
	}

	configPath := filepath.Join(user.HomeDir, ".wakatime.cfg")
	data, err := os.ReadFile(configPath)
	if !logError(err) {
		return
	}

	lines := strings.Split(string(data), "\n")

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
	var dailyStats []DailyStat

	httpResponse, err := sendRequest(config)
	if !logError(err) {
		return
	} else {
		defer httpResponse.Body.Close()
		x, err := io.ReadAll(httpResponse.Body)
		fmt.Println(string(x))
		if !logError(err) {
			return
		}
		response := Response{}
		err = json.Unmarshal(x, &response)
		if !logError(err) {
			return
		}
		fmt.Println("LANGUAGES:")

		for _, day := range response.Data {
			fmt.Println(day.Range.Date, day.Languages)
		}
		for _, day := range response.Data {
			dailyStats = append(dailyStats, DailyStat{
				Date:         day.Range.Date,
				TotalSeconds: day.GrandTotal.TotalSeconds,
				Projects:     day.Projects,
				Editors:      day.Editors,
				Languages:    day.Languages,
			})
		}

		graph(dailyStats)
		err = exportGraph(dailyStats, "sushupti.png", *darkBool)
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
