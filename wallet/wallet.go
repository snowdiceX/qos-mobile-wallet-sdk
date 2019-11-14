package wallet

import "github.com/QOSGroup/qos-mobile-wallet-sdk/account"

type Wallet struct {
	AM *account.AccountManager
}

func NewWallet(am *account.AccountManager) *Wallet  {
	return &Wallet{AM:am}
}



