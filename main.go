package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type urlStatus struct {
	url          string
	status       int
	errorMessage string
}

func sendRequest(url string, ch chan urlStatus) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		ch <- urlStatus{url, 0, err.Error()}
		return
	}

	ch <- urlStatus{url, res.StatusCode, ""}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go <url1> <url2> ... <urln>")
	}

	ch := make(chan urlStatus)

	for _, url := range os.Args[1:] {
		go sendRequest("https://"+url, ch)
	}

	results := make([]urlStatus, len(os.Args)-1)
	for i, _ := range results {
		results[i] = <-ch
		fmt.Printf("[%d] %s %s\n", results[i].status, results[i].url, results[i].errorMessage)
	}
}
