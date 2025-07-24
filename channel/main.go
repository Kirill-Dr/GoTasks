package main

import (
	"fmt"
	"net/http"
)

var GOOGLE_URL = "https://www.google.com"

func main() {
	code := make(chan int)
	go getHttpCode(code)
	<-code
}

func getHttpCode(codeCh chan int) {
	res, err := http.Get(GOOGLE_URL)
	if err != nil {
		fmt.Printf("Error fetching URL: %s\n", err.Error())
		return
	}
	fmt.Printf("HTTP Status Code: %d\n", res.StatusCode)
	codeCh <- res.StatusCode
}
