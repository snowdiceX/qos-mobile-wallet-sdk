package qos_mobile_wallet_sdk

import (
	"github.com/QOSGroup/qos-mobile-wallet-sdk/account"
	aApi "github.com/QOSGroup/qos-mobile-wallet-sdk/api"
	aWallet "github.com/QOSGroup/qos-mobile-wallet-sdk/wallet"
	db "github.com/tendermint/tm-db"

)

var api aApi.API

func InitWallet(storagePath string) {
	walletDB := db.NewDB("wallet", db.GoLevelDBBackend, storagePath)
	am := account.NewAccountManager(walletDB)
	api.Wallet = aWallet.NewWallet(am)
}
