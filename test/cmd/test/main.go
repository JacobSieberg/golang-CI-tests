package main

import (
	"fmt"
	"time"
)

const version = "5"

func main() {
	for {
		fmt.Printf("Version %s\n", version)
		time.Sleep(time.Second * 5)
	}
}
