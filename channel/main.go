package main

import (
	"fmt"
	"net/http"
	"sync"
)

var GOOGLE_URL = "https://www.google.com"

func main() {
	code := make(chan int)
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			getHttpCode(code)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(code)
	}()

	for res := range code {
		fmt.Printf("HTTP Status Code: %d\n", res)
	}
}

func getHttpCode(codeCh chan int) {
	res, err := http.Get(GOOGLE_URL)
	if err != nil {
		fmt.Printf("Error fetching URL: %s\n", err.Error())
		return
	}
	codeCh <- res.StatusCode
}
