package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/Dat-one-dev/Sushupti/data"
)

func LogError(err error) bool {
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}

func SendRequest(config data.Config) (*http.Response, error) {
	packet, err := http.NewRequest(http.MethodGet, config.APIUrl, nil)
	if !LogError(err) {
		return nil, err
	}

	packet.Header.Set("Authorization", "Bearer "+config.APIkey)

	packetClient := http.Client{}
	recievedReq, err := packetClient.Do(packet)
	if !LogError(err) {
		return nil, err
	}

	return recievedReq, nil
}

func ParseTime(date string) (string, error) {
	parsedTime, err := time.Parse("02-01-06", date)
	if err != nil {
		return "", err
	}

	return parsedTime.Format("2006-01-02"), nil
}

func LoadConfig(start, end string) (data.Config, error) {
	user, err := user.Current()
	if err != nil {
		return data.Config{}, err
	}
	configPath := filepath.Join(user.HomeDir, ".wakatime.cfg")
	configData, err := os.ReadFile(configPath)
	if err != nil {
		return data.Config{}, err
	}

	var config data.Config
	for _, line := range strings.Split(string(configData), "\n") {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "api_url":
			config.APIUrl = fmt.Sprintf(
				"%s/users/current/summaries?start=%s&end=%s",
				value,
				start,
				end,
			)
		case "api_key":
			config.APIkey = value
		}
	}
	return config, nil
}

func FetchDaily(config data.Config) ([]data.DailyStat, error) {
	response, err := SendRequest(config)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result data.Response

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	var daily []data.DailyStat

	for _, day := range result.Data {
		daily = append(daily, data.DailyStat{
			Date:         day.Range.Date,
			TotalSeconds: day.GrandTotal.TotalSeconds,
			Projects:     day.Projects,
			Editors:      day.Editors,
			Languages:    day.Languages,
		})
	}

	return daily, nil
}
