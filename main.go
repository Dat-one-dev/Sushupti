package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	configPath := filepath.Join(user.HomeDir, ".wakatime.cfg")
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(string(data))
}
