package main

import (
	"fmt"
)

func main() {
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter url: ")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Invalid format of data")
		return
	}
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
