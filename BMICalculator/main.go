package main

import (
	"fmt"
	"math"
)

func main() {
	const IMTPower = 2
	var userHeight float64
	var userWeight float64
	fmt.Print("___Calculator of BMI___\n")
	fmt.Print("Enter your height (m): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight (kg): ")
	fmt.Scan(&userWeight)
	IMT := userWeight / math.Pow(userHeight, IMTPower)
	fmt.Print("Your IMT is: ")
	fmt.Print(IMT)
}
