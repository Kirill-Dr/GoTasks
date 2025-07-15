package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var menuOperation = map[string]func([]float64) float64{
	"SUM": calculateSum,
	"AVG": calculateAvg,
	"MED": calculateMed,
}

var choiceOperation = []string{
	"SUM - Sum of numbers",
	"AVG - Average of numbers",
	"MED - Median of numbers",
	"EXIT - Exit program",
	"Choose an option (SUM/AVG/MED/EXIT)",
}

func main() {
	fmt.Println("--- Calculator ---")

	for {
		choice := promptData(choiceOperation...)
		choice = strings.TrimSpace(strings.ToUpper(choice))

		if choice == "EXIT" {
			break
		}

		operationFunc := menuOperation[choice]
		if operationFunc == nil {
			fmt.Println("Invalid choice. Exiting...")
			break
		}

		numbers := getNumbersFromUser()
		result := operationFunc(numbers)
		fmt.Printf("Result: %.2f\n", result)
	}
}

func promptData(prompt ...string) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
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

func calculateSum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calculateAvg(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func calculateMed(numbers []float64) float64 {
	sort.Float64s(numbers)
	mid := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2
	}
	return numbers[mid]
}
