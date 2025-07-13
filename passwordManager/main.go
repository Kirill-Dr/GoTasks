package main

import (
	"fmt"
	"passwordManager/account"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("--- Password Manager ---")
	vault := account.NewVault()
Menu:
	for {
		choice := getMenu()
		switch choice {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func createAccount(vault *account.Vault) {
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter url: ")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Invalid format of data")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := promptData("Enter url to find: ")
	foundAccounts := vault.FindAccountsByUrl(url)
	if len(foundAccounts) == 0 {
		color.Red("No accounts found")
		return
	}
	for index, account := range foundAccounts {
		color.Cyan("Account #" + strconv.Itoa(index+1))
		account.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Enter url to delete: ")
	deleted := vault.DeleteAccountByUrl(url)
	if deleted {
		color.Green("Account deleted")
	} else {
		color.Red("Account not found")
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
