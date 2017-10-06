package main

import (
	"fmt"

	"github.com/sanbornm/go-selfupdate/selfupdate"
)

func main() {
	var updater = &selfupdate.Updater{
		CurrentVersion: version,
		ApiURL:         "http://updates.yourdomain.com/",
		BinURL:         "http://updates.yourdomain.com/",
		DiffURL:        "http://updates.yourdomain.com/",
		Dir:            "update/",
		CmdName:        "myapp", // app name
	}

	if updater != nil {
		go updater.BackgroundRun()
	}
	fmt.Println("Hello World")
}
