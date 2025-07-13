package main

import (
	"fmt"
	"passwordManager/account"
	"passwordManager/files"
)

func main() {
	createAccount()
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
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Error marshalling into json:", err)
		return
	}
	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
