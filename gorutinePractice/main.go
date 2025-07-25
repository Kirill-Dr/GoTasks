package main

import (
	"fmt"
	"sync"
)

func sumNums(nums []int, ch chan int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	ch <- sum
}

func main() {
	var wg sync.WaitGroup
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 3
	ch := make(chan int, numGoroutines)

	partSize := len(arr) / numGoroutines

	for i := range numGoroutines {
		wg.Add(1)

		start := i * partSize
		end := start + partSize

		go func() {
			defer wg.Done()
			sumNums(arr[start:end], ch)
		}()
	}

	totalSum := 0
	for range numGoroutines {
		totalSum += <-ch
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Total Sum:", totalSum)
}
