package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("___Calculator of BMI___")
	userWeight, userHeight := getUserInput()
	BMI := calculateBMI(userWeight, userHeight)
	outputResult(BMI)
}

func outputResult(BMI float64) {
	result := fmt.Sprintf("Your BMI is: %.0f", BMI)
	fmt.Println(result)
}

func calculateBMI(userWeight float64, userHeight float64) (BMI float64) {
	BMI = userWeight / math.Pow(userHeight/100, IMTPower)
	return
}

func getUserInput() (float64, float64) {
	var userWeight float64
	var userHeight float64
	fmt.Print("Enter your weight (kg): ")
	fmt.Scan(&userWeight)
	fmt.Print("Enter your height (cm): ")
	fmt.Scan(&userHeight)
	return userWeight, userHeight
}
