package wallet

import (
	"errors"
	"github.com/QOSGroup/js-keys/keys"
	"github.com/QOSGroup/qos-mobile-wallet-sdk/account"
	"github.com/QOSGroup/qos-mobile-wallet-sdk/types"
	"github.com/tyler-smith/go-bip39"
	"strings"
)

const (
	defaultEntropySize = 256
)

type Wallet struct {
	am *account.AccountManager
}

func NewWallet(am *account.AccountManager) Wallet {
	return Wallet{am: am}
}

func (aWallet Wallet) GenerateMnemonic() Response {
	mnemonic, err := genMnemonic()
	if err != nil {
		return NewErrResponse(err)
	}
	return NewSuccessResponse(mnemonic)
}

func (aWallet Wallet) NewAccount(name, password, mnemonic string) Response {

	if len(password) == 0 {
		return NewErrResponse(errors.New("password not empty"))
	}

	if len(mnemonic) == 0 {
		mnemonic, _ = genMnemonic()
	}

	accMne, err := aWallet.am.NewAccountFromMnemonic(name, password, mnemonic)
	if err != nil {
		return NewErrResponse(err)
	}

	return NewSuccessResponse(accMne)
}

func (aWallet Wallet) ListAllAccounts() Response {
	return NewSuccessResponse(aWallet.am.ListAccounts())
}

func (aWallet Wallet) FindAccount(address string) Response {
	acc, err := aWallet.am.QueryAccount(address, "")
	if err != nil {
		return NewErrResponse(err)
	}
	return NewSuccessResponse(acc)
}

func (aWallet Wallet) FindAccountByName(name string) Response {
	acc, err := aWallet.am.QueryAccount("", name)
	if err != nil {
		return NewErrResponse(err)
	}
	return NewSuccessResponse(acc)
}

func (aWallet Wallet) ProbeAccount(address, password string) Response {
	acc, err := aWallet.am.QueryAccount(address, "")
	if err != nil {
		return NewErrResponse(err)
	}

	pk, err := aWallet.am.DecryptPrivateKey(acc, password)
	if err != nil {
		return NewErrResponse(err)
	}

	return NewSuccessResponse(account.NewAccountPrivacy(acc, pk))
}

func (aWallet Wallet) DeleteAccount(address, password string) Response {

	acc, err := aWallet.am.QueryAccount(address, "")
	if err != nil {
		return NewErrResponse(err)
	}

	_, err = aWallet.am.DecryptPrivateKey(acc, password)
	if err != nil {
		return NewErrResponse(err)
	}

	err = aWallet.am.DeleteAccount(acc, password)
	if err != nil {
		return NewErrResponse(err)
	}

	return NewSuccessResponse(nil)
}

func (aWallet Wallet) ModifyAccountPassword(address, oldPassword, newPassword string) Response {

	acc, err := aWallet.am.QueryAccount(address, "")
	if err != nil {
		return NewErrResponse(err)
	}

	newAcc, err := aWallet.am.ModifyAccountPassword(acc, oldPassword, newPassword)
	if err != nil {
		return NewErrResponse(err)
	}

	return NewSuccessResponse(newAcc)
}

func (aWallet Wallet) RecoverAccountFromPrivateKey(hexPrivateKey, password string) Response {

	priKeyBz := types.MustDecodeHex(hexPrivateKey)

	acc, err := aWallet.am.ImportAccount(priKeyBz, password)
	if err != nil {
		return NewErrResponse(err)
	}

	return NewSuccessResponse(acc)
}

func (aWallet Wallet) RecoverAccountFromMnemonic(mnemonic, password string) Response {

	priKeyBz, _, err := keys.DeriveQOSKey(mnemonic)
	if err != nil {
		return NewErrResponse(err)
	}

	acc, err := aWallet.am.ImportAccount(priKeyBz, password)
	if err != nil {
		return NewErrResponse(err)
	}

	return NewSuccessResponse(acc)
}

func (aWallet Wallet) SignData(address, password string, needSignStr string, codeType string) Response {

	acc, err := aWallet.am.QueryAccount(address, "")
	if err != nil {
		return NewErrResponse(err)
	}

	hexPk, err := aWallet.am.DecryptPrivateKey(acc, password)
	if err != nil {
		return NewErrResponse(err)
	}

	var data []byte
	if strings.EqualFold("base64", strings.ToLower(codeType)) {
		data = keys.Sign(types.MustDecodeHex(hexPk), types.MustDecodeBase64(needSignStr))
	} else {
		data = keys.Sign(types.MustDecodeHex(hexPk), []byte(needSignStr))
	}

	return NewSuccessResponse(types.EncodeBase64(data))
}

func genMnemonic() (string, error) {
	bz, err := bip39.NewEntropy(defaultEntropySize)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(bz)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}
