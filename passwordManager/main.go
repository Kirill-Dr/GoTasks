package main

import (
	"fmt"
	"passwordManager/account"
	"passwordManager/files"
	"passwordManager/output"
	"strconv"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	fmt.Println("--- Password Manager ---")
	vault := account.NewVault(files.NewJsonDB("data.json"))
Menu:
	for {
		choice := promptData([]string{"1. Create account", "2. Find account", "3. Delete account", "4. Exit", "Choose an option"})
		menuFunc := menu[choice]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDB) {
	login := promptData([]string{"Enter login"})
	password := promptData([]string{"Enter password"})
	url := promptData([]string{"Enter url"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Invalid format of data")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDB) {
	url := promptData([]string{"Enter url to find"})
	foundAccounts := vault.FindAccountsByUrl(url)
	if len(foundAccounts) == 0 {
		output.PrintError("No accounts found")
		return
	}
	for index, account := range foundAccounts {
		color.Cyan("Account #" + strconv.Itoa(index+1))
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := promptData([]string{"Enter url to delete"})
	deleted := vault.DeleteAccountByUrl(url)
	if deleted {
		color.Green("Account deleted")
	} else {
		output.PrintError("Account not found")
	}
}

func promptData[T any](prompt []T) string {
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
