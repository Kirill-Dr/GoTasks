package main

import (
	"fmt"
	"passwordManager/account"
	"passwordManager/files"
)

func main() {
	files.WriteFile("Hello", "test.txt")
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter url: ")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Invalid format of data")
		return
	}
	myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scanln(&res)
	return res
}
