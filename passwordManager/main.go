package main

import (
	"fmt"
	"passwordManager/account"
	"passwordManager/files"
)

func main() {
	fmt.Println("--- Password Manager ---")
Menu:
	for {
		choice := getMenu()
		switch choice {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var choice int
	fmt.Println("Choose an option:")
	fmt.Println("1. Create account")
	fmt.Println("2. Find account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")
	fmt.Scan(&choice)
	return choice
}

func createAccount() {
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Invalid format of data")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Error marshalling into json:", err)
		return
	}
	files.WriteFile(data, "data.json")
}

func findAccount() {}

func deleteAccount() {}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
