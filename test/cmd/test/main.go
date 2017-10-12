package main

import (
	"fmt"
	"time"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
)

const version = "4"
const url = "https://jacobtestupdate.blob.core.windows.net/updates/test-windows.amd64.exe"

func main() {
	overseer.Run(overseer.Config{
		Program: start,
		Fetcher: &fetcher.HTTP{
			URL:      url,
			Interval: 1 * time.Second,
		},
	})
}

func start(overseer.State) {
	fmt.Println("Application started")
	for {
		fmt.Printf("Current version: %s\n", version)
		time.Sleep(time.Second * 2)
	}
}
