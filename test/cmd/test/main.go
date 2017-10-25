package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//Version number
const Version = "2"

func main() {
	for {
		fmt.Printf("Version %s\n", Version)
		time.Sleep(time.Second * 5)
	}
}

func exitOnSig() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	fmt.Println(<-sigs)
	os.Exit(0)
}
