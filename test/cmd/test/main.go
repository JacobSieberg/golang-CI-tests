package main

import (
	"fmt"
	"time"
)

const version = "6"

func main() {
	for {
		fmt.Printf("Version %s\n", version)
		time.Sleep(time.Second * 5)
	}
}
