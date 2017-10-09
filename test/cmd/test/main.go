package main

import (
	"fmt"
	"time"

	"github.com/sanbornm/go-selfupdate/selfupdate"
)

func main() {
	var updater = &selfupdate.Updater{
		CurrentVersion: "42",
		ApiURL:         "https://jacobtestupdate.blob.core.windows.net/updates/",
		BinURL:         "https://jacobtestupdate.blob.core.windows.net/updates/",
		DiffURL:        "https://jacobtestupdate.blob.core.windows.net/updates/",
		Dir:            "update/",
		CmdName:        "test-go-app", // app name
	}

	if updater != nil {
		go func() {
			if err := updater.BackgroundRun(); err != nil {
				fmt.Println(err)
			}
		}()
	}

	for {
		fmt.Println("Version 42!")
		time.Sleep(time.Second)
	}
}
