package main

import (
	"crypto"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/inconshreveable/go-update"
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
		currentSum, err := getCurrentSha256Sum()
		if err != nil {
			fmt.Println(err)
		}
		if currentSum != sum.sum {
			fmt.Println("Update needed.")
			file, err := fetchBinary(settings, sum.name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Downloaded")
			fmt.Println("Starting update...")
			err = updateWithChecksum(file, sum.sum)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Updated")
			file.Close()
			fmt.Println("Restarting")
			cmd := exec.Command(os.Args[0], os.Args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			err = cmd.Start()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Restarted")
			os.Exit(0)

		}
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

func fetchBinary(settings settings, name string) (*os.File, error) {
	url := fmt.Sprintf("%s/%s", settings.baseURL, name)
	return download(name, url)
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

func download(filepath string, url string) (*os.File, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	file.Close()
	return os.Open(filepath)
}

func getCurrentSha256Sum() (string, error) {
	f, err := os.Open(os.Args[0])
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func updateWithChecksum(binary io.Reader, hexChecksum string) error {
	checksum, err := hex.DecodeString(hexChecksum)
	if err != nil {
		return err
	}
	err = update.Apply(binary, update.Options{
		Hash:     crypto.SHA256, // this is the default, you don't need to specify it
		Checksum: checksum,
	})
	if err != nil {
		update.RollbackError(err)
	}
	return err
}
