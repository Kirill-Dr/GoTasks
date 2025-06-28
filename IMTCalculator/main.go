package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	userHeight := 1.8
	userWeight := 100.0
	IMT := userWeight / math.Pow(userHeight, IMTPower)
	fmt.Print(IMT)
}
