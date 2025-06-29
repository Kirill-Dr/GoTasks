package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("___Calculator of BMI___")
	for {
		userWeight, userHeight := getUserInput()
		BMI, err := calculateBMI(userWeight, userHeight)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outputResult(BMI)
		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}
}

func outputResult(BMI float64) {
	result := fmt.Sprintf("Your BMI is: %.0f", BMI)
	fmt.Println(result)
	switch {
	case BMI < 16:
		fmt.Println("You are severely underweight")
	case BMI < 18.5:
		fmt.Println("You are underweight")
	case BMI < 25:
		fmt.Println("You are normal weight")
	case BMI < 30:
		fmt.Println("You are overweight")
	default:
		fmt.Println("You are obese")
	}
}

func calculateBMI(userWeight float64, userHeight float64) (float64, error) {
	if userWeight <= 0 || userHeight <= 0 {
		return 0, errors.New("Weight and height must be greater than 0")
	}
	BMI := userWeight / math.Pow(userHeight/100, IMTPower)
	return BMI, nil
}

func getUserInput() (float64, float64) {
	var userWeight float64
	var userHeight float64
	fmt.Print("Enter your weight (kg): ")
	_, errWeight := fmt.Scan(&userWeight)
	if errWeight != nil {
		var discard string
		fmt.Scanln(&discard)
		return 0, 0
	}
	fmt.Print("Enter your height (cm): ")
	_, errHeight := fmt.Scan(&userHeight)
	if errHeight != nil {
		var discard string
		fmt.Scanln(&discard)
		return 0, 0
	}
	return userWeight, userHeight
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Do you want to continue? (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false
}
