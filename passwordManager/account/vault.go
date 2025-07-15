package account

import (
	"encoding/json"
	"passwordManager/encrypter"
	"passwordManager/output"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type DB interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDB struct {
	Vault
	db  DB
	enc encrypter.Encrypter
}

func NewVault(db DB, enc encrypter.Encrypter) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	decryptedData := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(decryptedData, &vault)
	color.Cyan("Found %d accounts", len(vault.Accounts))
	if err != nil {
		output.PrintError(err)
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDB) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.save()
}

func (vault *VaultWithDB) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var foundAccounts []Account
	for _, account := range vault.Accounts {
		isMatched := checker(account, str)
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
		output.PrintError(err)
	}
	encryptedData := vault.enc.Encrypt(data)
	vault.db.Write(encryptedData)
}
