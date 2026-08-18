package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	user, err := user.Current()
	logError(err)

	configPath := filepath.Join(user.HomeDir, ".wakatime.cfg")
	data, err := os.ReadFile(configPath)
	logError(err)

	fmt.Println(string(data))
}

func logError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
