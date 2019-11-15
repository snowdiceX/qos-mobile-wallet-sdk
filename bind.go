package qos_mobile_wallet_sdk

import (
	"encoding/json"
	"github.com/QOSGroup/qos-mobile-wallet-sdk/account"
	aWallet "github.com/QOSGroup/qos-mobile-wallet-sdk/wallet"
	db "github.com/tendermint/tm-db"
)

var _wallet aWallet.Wallet

func InitWallet(storagePath string) {
	walletDB := db.NewDB("wallet", db.GoLevelDBBackend, storagePath)
	am := account.NewAccountManager(walletDB)
	_wallet = aWallet.NewWallet(am)
}

func ProduceMnemonic() string {
	str, _ := json.Marshal(_wallet.GenerateMnemonic())
	return string(str)
}

func CreateAccount(password string) string {
	str, _ := json.Marshal(_wallet.NewAccount("", password, ""))
	return string(str)
}

func CreateAccountWithName(name, password string) string {
	str, _ := json.Marshal(_wallet.NewAccount(name, password, ""))
	return string(str)
}

func CreateAccountWithMnemonic(name, password, mnemonic string) string {
	str, _ := json.Marshal(_wallet.NewAccount(name, password, mnemonic))
	return string(str)
}

func GetAccount(address string) string {
	str, _ := json.Marshal(_wallet.FindAccount(address))
	return string(str)
}

func GetAccountByName(name string) string {
	str, _ := json.Marshal(_wallet.FindAccountByName(name))
	return string(str)
}

func DeleteAccount(address, password string) string {
	str, _ := json.Marshal(_wallet.DeleteAccount(address, password))
	return string(str)
}

func ExportAccount(address, password string) string {
	str, _ := json.Marshal(_wallet.ProbeAccount(address, password))
	return string(str)
}

func ImportMnemonic(mnemonic, password string) string {
	str, _ := json.Marshal(_wallet.RecoverAccountFromMnemonic(mnemonic, password))
	return string(str)
}

func ImportPrivateKey(hexPrivateKey, password string) string {
	str, _ := json.Marshal(_wallet.RecoverAccountFromPrivateKey(hexPrivateKey, password))
	return string(str)
}

func ListAllAccounts() string {
	str, _ := json.Marshal(_wallet.ListAllAccounts())
	return string(str)
}

func Sign(address, password, signStr string) string {
	str, _ := json.Marshal(_wallet.SignData(address, password, signStr, ""))
	return string(str)
}

func SignBase64(address, password, base64Str string) string {
	str, _ := json.Marshal(_wallet.SignData(address, password, base64Str, "base64"))
	return string(str)
}
