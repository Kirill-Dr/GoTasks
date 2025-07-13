package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type Account struct {
	login    string
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letters[rand.IntN(len(letters))]
	}
	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}

	if password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}
