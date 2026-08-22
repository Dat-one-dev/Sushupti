package main

import (
	"fmt"
	"net/http"
	"time"
)

func logError(err error) bool {
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}

func sendRequest(config Config) (*http.Response, error) {
	packet, err := http.NewRequest(http.MethodGet, config.APIUrl, nil)
	if !logError(err) {
		return nil, err
	}

	packet.Header.Set("Authorization", "Bearer "+config.APIkey)

	packetClient := http.Client{}
	recievedReq, err := packetClient.Do(packet)
	if !logError(err) {
		return nil, err
	}

	return recievedReq, nil
}

func parseTime(date string) (string, error) {
	parsedTime, err := time.Parse("02-01-06", date)
	if err != nil {
		return "", err
	}

	return parsedTime.Format("2006-01-02"), nil
}

func mostUsedProject(daily []DailyStat) Project {
	var most Project

	for _, day := range daily {
		for _, project := range day.Projects {
			if project.TotalSeconds > most.TotalSeconds {
				most = project
			}
		}
	}

	return most
}

func mostUsedEditor(daily []DailyStat) Category {
	var most Category

	for _, day := range daily {
		for _, editor := range day.Editors {
			if editor.TotalSeconds > most.TotalSeconds {
				most = editor
			}
		}
	}

	return most
}

func mostUsedlang(daily []DailyStat) Category {
	var most Category

	for _, day := range daily {
		for _, lang := range day.Languages {
			if lang.TotalSeconds > most.TotalSeconds {
				most = lang
			}
		}
	}

	return most
}

func projectTotalTime(daily []DailyStat, projectName string) int {
	total := 0
	for _, day := range daily {
		for _, project := range day.Projects {
			if project.ProjectName == projectName {
				total += project.TotalSeconds
			}
		}
	}

	return total
}

func projectShareCalc(daily []DailyStat, projectTime int) float64 {
	total := 0
	for _, day := range daily {
		total += day.TotalSeconds
	}

	share := 0.0

	if total > 0 {
		share = float64(projectTime) / float64(total) * 100
	}

	return share
}