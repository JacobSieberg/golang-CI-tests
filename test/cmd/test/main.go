package main

import (
	"fmt"
	"os"
	"time"
)

const version = "12"

func main() {
	go func() {
		time.Sleep(time.Minute)
		os.Exit(0)
	}()
	for {
		fmt.Printf("Version %s\n", version)
		time.Sleep(time.Second * 5)
	}
}
