package main

import (
	"fmt"
)

type exchangeRatesType = map[string]map[string]float64

func main() {
	var exchangeRates = exchangeRatesType{
		"USD": {
			"EUR": 0.85,
			"RUB": 78.08,
		},
		"EUR": {
			"USD": 1.17,
			"RUB": 91.45,
		},
		"RUB": {
			"USD": 0.013,
			"EUR": 0.011,
		},
	}

	var currencyOrder = []string{"USD", "EUR", "RUB"}

	fmt.Println("--- Converter currencies ---")

	for {
		originalCurrency, amountOfMoney, targetCurrency := getUserInput(&currencyOrder)
		result := calculate(originalCurrency, amountOfMoney, targetCurrency, &exchangeRates)
		fmt.Printf("%0.2f %s = %0.2f %s", amountOfMoney, originalCurrency, result, targetCurrency)
		shouldContinue := askToContinue()
		if !shouldContinue {
			break
		}
	}
}

func getAvailableCurrencies(currencyOrder *[]string, excludeCurrency string) []string {
	var availableCurrencies []string
	for _, currency := range *currencyOrder {
		if currency != excludeCurrency {
			availableCurrencies = append(availableCurrencies, currency)
		}
	}
	return availableCurrencies
}

func formatCurrencies(currencies []string) string {
	result := ""
	for i, currency := range currencies {
		if i > 0 {
			result += " "
		}
		result += currency
	}
	return result
}

func getUserInput(currencyOrder *[]string) (string, float64, string) {
	var originalCurrency string
	var amountOfMoney float64
	var targetCurrency string

	for {
		availableCurrencies := getAvailableCurrencies(currencyOrder, "")
		fmt.Print("Enter original currency (" + formatCurrencies(availableCurrencies) + "): ")
		fmt.Scan(&originalCurrency)

		validCurrency := false
		for _, currency := range availableCurrencies {
			if currency == originalCurrency {
				validCurrency = true
				break
			}
		}
		if !validCurrency {
			fmt.Println("Invalid currency. Available currencies:", formatCurrencies(availableCurrencies))
			continue
		}
		break
	}

	for {
		fmt.Print("Enter amount of money: ")
		fmt.Scan(&amountOfMoney)
		if amountOfMoney <= 0 {
			fmt.Println("Amount must be greater than 0")
			continue
		}
		break
	}

	for {
		availableCurrencies := getAvailableCurrencies(currencyOrder, originalCurrency)
		fmt.Print("Enter target currency (" + formatCurrencies(availableCurrencies) + "): ")
		fmt.Scan(&targetCurrency)

		validCurrency := false
		for _, currency := range availableCurrencies {
			if currency == targetCurrency {
				validCurrency = true
				break
			}
		}
		if !validCurrency {
			fmt.Println("Invalid currency. Available currencies:", formatCurrencies(availableCurrencies))
			continue
		}
		break
	}

	return originalCurrency, amountOfMoney, targetCurrency
}

func calculate(originalCurrency string, amountOfMoney float64, targetCurrency string, exchangeRates *exchangeRatesType) float64 {
	rate := (*exchangeRates)[originalCurrency][targetCurrency]
	return amountOfMoney * rate
}

func askToContinue() bool {
	var answer string
	for {
		fmt.Print("\nDo you want to continue? (y/n): ")
		fmt.Scan(&answer)
		if answer == "y" || answer == "Y" {
			return true
		} else if answer == "n" || answer == "N" {
			return false
		} else {
			fmt.Println("Invalid input. Please enter y or n")
		}
	}
}
