package main

import (
	"fmt"
	"os/user"
	"path/filepath"
)

func main() {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	
}
