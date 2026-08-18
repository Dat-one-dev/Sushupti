package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Config struct {
	APIUrl string
	APIkey string
}

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

func main() {
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
			config.APIUrl = value + "/users/current/summaries?start=2026-08-18&end=2026-08-19"
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
		} else {
			fmt.Println("Body: ", string(x))
		}
	}

}
