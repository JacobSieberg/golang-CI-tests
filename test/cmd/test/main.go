package main

import (
	"fmt"
	"time"

	"github.com/sanbornm/go-selfupdate/selfupdate"
)

const version = "50"

func main() {
	var updater = &selfupdate.Updater{
		CurrentVersion: version,
		ApiURL:         "https://jacobtestupdate.blob.core.windows.net/updates/",
		BinURL:         "https://jacobtestupdate.blob.core.windows.net/updates/",
		DiffURL:        "https://jacobtestupdate.blob.core.windows.net/updates/",
		Dir:            "update/",
		CmdName:        "testapp", // app name
	}

	if updater != nil {
		go func() {
			if err := updater.BackgroundRun(); err != nil {
				fmt.Println(err)
			}
		}()
	}

	for {
		fmt.Printf("Version %s!\n", version)
		time.Sleep(time.Second)
	}
}
