package main

import (
	"errors"
	"fmt"
)

func main() {
	transactions := []float64{}
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
	fmt.Println(transactions)
}

func scanTransaction() (float64, error) {
	var transaction float64
	fmt.Println("Enter transaction (n to finish): ")
	_, err := fmt.Scan(&transaction)
	if err != nil {
		var discard string
		fmt.Scanln(&discard)
		return 0, errors.New("Transaction is not valid")
	}
	return transaction, nil
}
