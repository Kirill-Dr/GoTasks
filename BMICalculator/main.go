package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Println("___Calculator of BMI___")
	fmt.Print("Enter your height (cm): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight (kg): ")
	fmt.Scan(&userWeight)
	BMI := userWeight / math.Pow(userHeight/100, IMTPower)
	outputResult(BMI)
}

func outputResult(BMI float64) {
	result := fmt.Sprintf("Your BMI is: %.0f", BMI)
	fmt.Println(result)
}
