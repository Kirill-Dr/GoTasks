package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter url: ")

	account := account{
		password: password,
		url:      url,
		login:    login,
	}

	outputPassword(login, password, url)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(login, password, url string) {
	fmt.Println(login, password, url)
}
