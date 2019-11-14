package account

import (
	db "github.com/tendermint/tm-db"
	"github.com/tyler-smith/go-bip39"
	"sync"
)
import "github.com/QOSGroup/js-keys/keys"

const (
	defaultEntropySize = 256
)

type Account struct {
	Id uint
	Name string
	PubKey string
	Address string
	Mnemonic string
}

type AccountManager struct {
	db db.DB
	lock sync.Mutex
}

func GenerateMnemonic() (string, error) {
	bz, err := bip39.NewEntropy(defaultEntropySize)
	if err != nil {
		return "", err
	}

	return bip39.NewMnemonic(bz)
}

func ParseMnemonic(mnemonic string) ([]byte, error) {
	priKey, _, err := keys.DeriveQOSKey(mnemonic)
	return priKey, err
}


func NewAccountManager(db db.DB) *AccountManager {
	return &AccountManager{
		db: db,
	}
}

func (manager *AccountManager) CreateAccount()  {

}

func (manager *AccountManager) CreateAccountFromMnemonic()  {

}

func (manager *AccountManager) FindAccount()  {

}

func (manager *AccountManager) ListAccounts()  {

}

func (manager *AccountManager) DeleteAccount()  {

}






