package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}

func main() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Server is running")
	http.ListenAndServe(":8081", nil)
}
