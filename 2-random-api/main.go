package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

func randomNumber(w http.ResponseWriter, req *http.Request) {
	num := rand.Intn(6) + 1

	byteNum := strconv.Itoa(num)
	w.Write([]byte(byteNum))
}

func main() {
	http.HandleFunc("/random", randomNumber)

	http.ListenAndServe(":8081", nil)
}
