package main

import (
	"fmt"
	"time"
)

const version = "50"
const url = "https://jacobtestupdate.blob.core.windows.net/updates/"

func main() {
	go waitForUpdates(settings{baseURL: url, checkInterval: time.Second * 5})
	for {
		fmt.Printf("Current version: %s\n", version)
		time.Sleep(time.Second * 2)
	}
}
