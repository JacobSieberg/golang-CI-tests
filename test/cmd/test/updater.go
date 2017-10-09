package main

import (
	"fmt"
	"time"
)

type settings struct {
	baseURL       string
	checkInterval time.Duration
}

func waitForUpdates(settings settings) {
	for {
		fmt.Println("Checking for updates.")
		time.Sleep(settings.checkInterval)
	}
}
