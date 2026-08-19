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
