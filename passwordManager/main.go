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
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &account{
		url:      urlString,
		login:    login,
		password: password,
	}

	if password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main() {
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter url: ")

	myAccount, err := newAccount(login, password, url)
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
