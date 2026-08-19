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
)

func main() {
	startDate := flag.String("s", "", "start date")
	endDate := flag.String("e", "", "end date")
	flag.Parse()
	fmt.Println("Start:", *startDate)
	fmt.Println("End:", *endDate)
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
			config.APIUrl = value + "/users/current/summaries?start=2026-08-01&end=2026-08-18"
		}

		if key == "api_key" {
			config.APIkey = value
		}
	}

	httpResponse, err := sendRequest(config)
	if !logError(err) {
		return
	} else {
		defer httpResponse.Body.Close()
		fmt.Println("Status:", httpResponse.Status)
		fmt.Println("Status code:", httpResponse.StatusCode)
		x, err := io.ReadAll(httpResponse.Body)
		if !logError(err) {
			return
		}
		response := Response{}
		err = json.Unmarshal(x, &response)
		if !logError(err) {
			return
		}

		var dailyStats []DailyStat

		for _, day := range response.Data {
			dailyStats = append(dailyStats, DailyStat{
				Date:         day.Range.Date,
				TotalSeconds: day.GrandTotal.TotalSeconds,
			})
		}

		graph(dailyStats)

	}

}
