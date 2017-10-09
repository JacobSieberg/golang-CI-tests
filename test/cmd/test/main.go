package main

import "time"

const version = "50"
const url = "https://jacobtestupdate.blob.core.windows.net/updates/"

func main() {
	go waitForUpdates(settings{baseURL: url, checkInterval: time.Second * 5})
	for {
	}
}
