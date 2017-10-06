package main

import (
	"fmt"

	"github.com/sanbornm/go-selfupdate/selfupdate"
)

func main() {
	var updater = &selfupdate.Updater{
		CurrentVersion: "1.2",
		ApiURL:         "http://updates.yourdomain.com/",
		BinURL:         "http://updates.yourdomain.com/",
		DiffURL:        "http://updates.yourdomain.com/",
		Dir:            "update/",
		CmdName:        "myapp", // app name
	}

	if updater != nil {
		go func() {
			if err := updater.BackgroundRun(); err != nil {
				fmt.Println(err)
			}
		}()
	}
	fmt.Println("Hello World")
}
