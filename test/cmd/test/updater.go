package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

type settings struct {
	baseURL       string
	checkInterval time.Duration
}

type sha256Sum struct {
	sum  string
	name string
}

func waitForUpdates(settings settings) {
	for {
		fmt.Println("Checking for updates.")
		sum, err := fetchCheckSum(settings)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(sum.name)
		fmt.Println(sum.sum)
		time.Sleep(settings.checkInterval)
	}
}

func fetchCheckSum(settings settings) (sha256Sum, error) {
	url := fmt.Sprintf("%s/%s-%s.sum", settings.baseURL, runtime.GOOS, runtime.GOARCH)
	bytes, err := get(url)
	if err != nil {
		return sha256Sum{}, err
	}
	if len(bytes) < 67 {
		return sha256Sum{}, fmt.Errorf("expected 67 or more bytes, got %v", len(bytes))
	}
	return sha256Sum{
		name: string(bytes[66 : len(bytes)-1]),
		sum:  string(bytes[0:64]),
	}, nil
}

func get(url string) ([]byte, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := netClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body, err
}
