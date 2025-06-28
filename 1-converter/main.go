package main

import "fmt"

const USDToEUR float64 = 0.85
const USDToRUB float64 = 78.22

func main() {
	amount := getUserInput()
	result := calculate(amount, "EUR", "RUB")
	fmt.Printf("%.2f EUR = %.2f RUB.", amount, result)
}

func getUserInput() float64 {
	var amount float64
	fmt.Print("Enter amount of money: ")
	fmt.Scan(&amount)
	return amount
}

func calculate(amount float64, originalCurrency string, targetCurrenct string) float64 {
	EURToRUB := (amount / USDToEUR) * USDToRUB
	return EURToRUB
}
