package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- Calculator ---")
	operation := chooseOperation()
	numbers := getNumbersFromUser()
	result := calculate(operation, numbers)
	fmt.Printf("Result: %.2f\n", result)
}

func chooseOperation() string {
	for {
		fmt.Println("Choose operation:")
		fmt.Println("1. SUM")
		fmt.Println("2. AVG")
		fmt.Println("3. MED")
		fmt.Print("Enter your choice (SUM/AVG/MED):")

		var choice string
		_, err := fmt.Scanln(&choice)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Error: input is not valid")
			fmt.Println("Invalid choice. Please enter SUM, AVG, or MED.")
			continue
		}

		choice = strings.TrimSpace(strings.ToUpper(choice))

		if choice == "SUM" || choice == "AVG" || choice == "MED" {
			return choice
		}

		fmt.Println("Invalid choice. Please enter SUM, AVG, or MED.")
	}
}

func getNumbersFromUser() []float64 {
	for {
		fmt.Print("Enter numbers separated by commas (e.g. 1,2,3): ")

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			var discard string
			fmt.Scanln(&discard)
			fmt.Println("Error: input is not valid")
			fmt.Println("Please try again with correct format (e.g. 1,2,3):")
			continue
		}

		if !strings.Contains(input, ",") {
			fmt.Println("Error: numbers must be separated by commas without spaces")
			fmt.Println("Please try again with correct format (e.g. 1,2,3):")
			continue
		}

		parts := strings.Split(input, ",")

		var numbers []float64
		allValid := true

		for _, part := range parts {
			part = strings.TrimSpace(part)

			if part == "" {
				continue
			}

			number, err := strconv.ParseFloat(part, 64)
			if err != nil {
				fmt.Println("Error: not a valid number")
				allValid = false
				break
			}

			numbers = append(numbers, number)
		}

		if allValid && len(numbers) > 0 {
			return numbers
		}

		fmt.Println("Please try again with correct format (e.g. 1,2,3)")
	}
}

func calculate(operation string, numbers []float64) float64 {
	switch operation {
	case "SUM":
		sum := 0.0
		for _, num := range numbers {
			sum += num
		}
		return sum

	case "AVG":
		sum := 0.0
		for _, num := range numbers {
			sum += num
		}
		return sum / float64(len(numbers))

	case "MED":
		sort.Float64s(numbers)
		mid := len(numbers) / 2
		if len(numbers)%2 == 0 {
			return (numbers[mid-1] + numbers[mid]) / 2
		}
		return numbers[mid]

	default:
		return 0
	}
}
