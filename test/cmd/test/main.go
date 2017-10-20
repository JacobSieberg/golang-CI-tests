package main

import (
	"fmt"
	"time"
)

const version = "11"

func main() {
	for {
		fmt.Printf("Version %s\n", version)
		time.Sleep(time.Second * 5)
	}
}
