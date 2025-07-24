package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const GOOGLE_URL = "https://www.google.com"

func main() {
	t := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHttpCode()
			wg.Done()
		}()
	}

	fmt.Println(time.Since(t))
	wg.Wait()
	fmt.Println("Finished fetching HTTP codes")
}

func getHttpCode() {
	res, err := http.Get(GOOGLE_URL)
	if err != nil {
		fmt.Printf("Error fetching URL: %s\n", err.Error())
		return
	}
	fmt.Printf("HTTP Status Code: %d\n", res.StatusCode)
}
