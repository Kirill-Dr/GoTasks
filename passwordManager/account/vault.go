package account

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type DB interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDB struct {
	Vault
	db DB
}

func NewVault(db DB) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Error parsing data.json")
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDB) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.save()
}

func (vault *VaultWithDB) FindAccountsByUrl(url string) []Account {
	var foundAccounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if isMatched {
			foundAccounts = append(foundAccounts, account)
		}
	}
	return foundAccounts
}

func (vault *VaultWithDB) DeleteAccountByUrl(url string) bool {
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

func (vault *VaultWithDB) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.Vault.ToBytes()
	if err != nil {
		color.Red("Error serializing data.json")
	}
	vault.db.Write(data)
}
