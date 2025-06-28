package main

import "fmt"

func main() {
	const USDToEUR float64 = 0.85
	const USDToRUB float64 = 78.22

	EURToRUB := (1.0 / USDToEUR) * USDToRUB

	fmt.Printf("1 EUR = %.2f RUB.", EURToRUB)
}
