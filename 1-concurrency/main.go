package main

import (
	"fmt"
	"math/rand"
)

func genRandomNumber(numCh chan int) {
	defer close(numCh)
	for range 10 {
		numCh <- rand.Intn(101)
	}
}

func square(numCh <-chan int, squareCh chan int) {
	defer close(squareCh)
	for num := range numCh {
		squareCh <- num * num
	}
}

func main() {
	numCh := make(chan int)
	squareCh := make(chan int)

	numSlice := make([]int, 10)

	go genRandomNumber(numCh)
	go square(numCh, squareCh)

	for i := range numSlice {
		numSlice[i] = <-squareCh
	}

	fmt.Println("Squared numbers:", numSlice)
}
