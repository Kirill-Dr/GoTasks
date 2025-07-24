package main

import (
	"fmt"
	"net/http"
	"time"
)

const GOOGLE_URL = "https://www.google.com"

func main() {
	t := time.Now()
	for i := 0; i < 10; i++ {
		go getHttpCode()
	}

	fmt.Println(time.Since(t))
	time.Sleep(time.Second * 3)
	fmt.Println("Finished fetching HTTP codes")
}

func getHttpCode() {
	res, err := http.Get(GOOGLE_URL)
	if err != nil {
		fmt.Printf("Error fetching URL: %s\n", err.Error())
	}
	fmt.Printf("HTTP Status Code: %d\n", res.StatusCode)
}
