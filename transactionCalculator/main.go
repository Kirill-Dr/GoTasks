package main

import (
	"errors"
	"fmt"
)

func main() {
	transactions := []float64{}
	fmt.Println("--- Transaction Calculator ---")
	for {
		transaction, err := scanTransaction()
		if transaction == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		transactions = append(transactions, transaction)
	}
	balance := calculateBalance(transactions)
	fmt.Printf("Your balance is: %.2f\n", balance)
}

func scanTransaction() (float64, error) {
	var transaction float64
	fmt.Println("Enter transaction (0 to finish): ")
	_, err := fmt.Scan(&transaction)
	if err != nil {
		var discard string
		fmt.Scanln(&discard)
		return 0, errors.New("Transaction is not valid")
	}
	return transaction, nil
}

func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance += value
	}
	return balance
}
