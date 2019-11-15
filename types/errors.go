package types

import "errors"

var (
	AccountExistsError    = errors.New("account has already exists")
	AccountNotExistsError = errors.New("account not exists")
	PasswordError         = errors.New("password not right")
)
