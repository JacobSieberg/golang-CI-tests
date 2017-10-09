package main

import (
	"fmt"
	"time"
)

const version = "4"
const url = "https://jacobtestupdate.blob.core.windows.net/updates/"

func main() {
	go waitForUpdates(settings{baseURL: url, checkInterval: time.Second * 5})
	fmt.Println("Application started")
	for {
		fmt.Printf("Current version: %s\n", version)
		time.Sleep(time.Second * 2)
	}
}
