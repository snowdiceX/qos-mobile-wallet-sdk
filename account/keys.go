package account

import "encoding/binary"

var (
	accountIndexKey = []byte("AccountIndex")
	accountPrefix = []byte("account:")
	accountNamePrefixKey = []byte("account-name:")
)

func AccountIndexKey() []byte {
	return accountIndexKey
}

func AccountKey(index uint64) []byte {
	bz := make([]byte, 0, 8)
	binary.BigEndian.PutUint64(bz, index)
	return append(accountPrefix, bz...)
}

func AccountNameIndexKey(name string) []byte {
	return append(accountNamePrefixKey, []byte(name)...)
}