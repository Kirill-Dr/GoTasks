package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letters[rand.IntN(len(letters))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	return &account{
		url:      urlString,
		login:    login,
		password: password,
	}, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main() {
	login := promptData("Enter login: ")
	url := promptData("Enter url: ")

	myAccount, err := newAccount(login, "", url)
	if err != nil {
		fmt.Println("Invalid format of URL")
		return
	}
	myAccount.generatePassword(10)
	myAccount.outputPassword()
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}
