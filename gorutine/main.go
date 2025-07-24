package main

import (
	"fmt"
	"time"
)

func main() {
	go printHi()
	fmt.Println("Hello main")
	time.Sleep(time.Second)
}

func printHi() {
	fmt.Println("Hello gorutine")
}
