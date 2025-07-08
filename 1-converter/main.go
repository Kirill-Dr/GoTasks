package main

import (
	"errors"
	"fmt"
)

var exchangeRates = map[string]map[string]float64{
	"USD": {
		"EUR": 0.85,
		"RUB": 78.22,
	},
	"EUR": {
		"USD": 1.0 / 0.85,
		"RUB": 78.22 / 0.85,
	},
	"RUB": {
		"USD": 1.0 / 78.22,
		"EUR": 0.85 / 78.22,
	},
}

var currencyOrder = []string{"USD", "EUR", "RUB"}

func main() {
	fmt.Println("--- Converter currencies ---")

	for {
		originalCurrency, amount, targetCurrency := getUserInput()
		result := calculate(amount, originalCurrency, targetCurrency)
		fmt.Printf("%.2f %s = %.2f %s\n", amount, originalCurrency, result, targetCurrency)

		if !askToContinue() {
			break
		}
	}
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

func getAvailableCurrencies(excludeCurrency string) []string {
	var available []string
	for currency := range exchangeRates {
		if currency != excludeCurrency {
			available = append(available, currency)
		}
	}
	return available
}

func getAllCurrencies() string {
	return formatCurrencies(currencyOrder)
}

func getUserInput() (string, float64, string) {
	var originalCurrency string
	var amount float64
	var targetCurrency string

	for {
		fmt.Print("Enter original currency " + getAllCurrencies() + ": ")
		_, errorOriginalCurrency := fmt.Scan(&originalCurrency)
		if errorOriginalCurrency != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Original currency is not valid " + getAllCurrencies())
			continue
		}

		if _, exists := exchangeRates[originalCurrency]; !exists {
			fmt.Println("Currency '" + originalCurrency + "' is not supported. Available: " + getAllCurrencies())
			continue
		}
		break
	}

	for {
		fmt.Print("Enter amount of money: ")
		_, errorAmount := fmt.Scan(&amount)
		if amount <= 0 {
			fmt.Println("Amount must be greater than 0")
			continue
		}
		if errorAmount != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Amount is not valid")
			continue
		}
		break
	}

	availableCurrencies := getAvailableCurrencies(originalCurrency)
	formattedCurrencies := formatCurrencies(availableCurrencies)

	for {
		fmt.Printf("Enter target currency (%s): ", formattedCurrencies)
		_, errorTargetCurrency := fmt.Scan(&targetCurrency)
		if errorTargetCurrency != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Target currency is not valid")
			continue
		}
		if _, exists := exchangeRates[targetCurrency]; !exists {
			availableCurrencies := getAvailableCurrencies(originalCurrency)
			formattedAvailable := formatCurrencies(availableCurrencies)
			fmt.Println("Currency '" + targetCurrency + "' is not supported. Available: " + formattedAvailable)
			continue
		}
		if originalCurrency == targetCurrency {
			fmt.Println("Original and target currencies cannot be the same")
			continue
		}
		break
	}

	return originalCurrency, amount, targetCurrency
}

func getExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	if rates, exists := exchangeRates[fromCurrency]; exists {
		if rate, exists := rates[toCurrency]; exists {
			return rate, nil
		}
	}

	return 0, errors.New("Exchange rate not found")
}

func calculate(amount float64, originalCurrency string, targetCurrency string) float64 {
	rate, _ := getExchangeRate(originalCurrency, targetCurrency)
	return amount * rate
}

func askToContinue() bool {
	var choice string
	for {
		fmt.Print("Do you want to continue? (y/n): ")
		fmt.Scan(&choice)
		if choice == "y" || choice == "Y" {
			return true
		}
		return false
	}
}
