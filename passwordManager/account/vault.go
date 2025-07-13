package account

import (
	"encoding/json"
	"passwordManager/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewVault() *Vault {
	db := files.NewJsonDB("data.json")
	file, err := db.Read()
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Error parsing data.json")
	}
	return &vault
}

func (vault *Vault) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.save()
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var foundAccounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			foundAccounts = append(foundAccounts, account)
		}
	}
	return foundAccounts
}

func (vault *Vault) DeleteAccountByUrl(url string) bool {
	for i := 0; i < len(vault.Accounts); i++ {
		if vault.Accounts[i].Url == url {
			vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...)
			vault.save()

			return true
		}
	}
	return false
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Error serializing data.json")
	}
	db := files.NewJsonDB("data.json")
	db.Write(data)
}
