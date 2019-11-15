package types

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
)

var (
	commonIV                  = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	qosVerificationStartBytes = []byte("-- QOS Wallet SDK Start Verification --")
	qosVerificationEndBytes   = []byte("-- QOS Wallet SDK End Verification --")
)

func Encrypt(data []byte, passphrase string) (encBytes []byte, saltBytes []byte) {
	saltBytes = randBytes(16)
	key := hash(append([]byte(passphrase), saltBytes...))

	block, _ := aes.NewCipher(key)
	cfb := cipher.NewCFBEncrypter(block, commonIV)

	copyData := make([]byte, len(data)+len(qosVerificationStartBytes)+len(qosVerificationEndBytes))
	copy(copyData[0:len(qosVerificationStartBytes)], qosVerificationStartBytes)
	copy(copyData[len(qosVerificationStartBytes):len(data)+len(qosVerificationStartBytes)], data)
	copy(copyData[len(data)+len(qosVerificationStartBytes):], qosVerificationEndBytes)

	encBytes = make([]byte, len(copyData))
	cfb.XORKeyStream(encBytes, copyData)
	return
}

func Decrypt(encBytes []byte, saltBytes []byte, passphrase string) ([]byte, error) {
	key := hash(append([]byte(passphrase), saltBytes...))

	block, _ := aes.NewCipher(key)
	cfb := cipher.NewCFBDecrypter(block, commonIV)

	data := make([]byte, len(encBytes))
	cfb.XORKeyStream(data, encBytes)

	startBz := data[0:len(qosVerificationStartBytes)]
	endBz := data[len(data)-len(qosVerificationEndBytes):]

	if bytes.Equal(startBz, qosVerificationStartBytes) && bytes.Equal(endBz, qosVerificationEndBytes) {
		return data[len(qosVerificationStartBytes) : len(data)-len(qosVerificationEndBytes)], nil
	}

	return nil, PasswordError
}

func hash(bz []byte) []byte {
	hash := sha256.New()
	hash.Write(bz)
	return hash.Sum(nil)
}

func EncodeBase64(pk []byte) string {
	return base64.StdEncoding.EncodeToString(pk)
}

func MustDecodeBase64(encPK string) []byte {
	bz, err := base64.StdEncoding.DecodeString(encPK)
	if err != nil {
		panic(err)
	}
	return bz
}

func EncodeHex(pk []byte) string {
	return hex.EncodeToString(pk)
}

func MustDecodeHex(encPK string) []byte {
	bz, err := hex.DecodeString(encPK)
	if err != nil {
		panic(err)
	}
	return bz
}

func BytesToUint64(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func Uint64ToBytes(u uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, u)
	return bz
}

func PrefixEndBytes(prefix []byte) []byte {
	if len(prefix) == 0 {
		return nil
	}

	end := make([]byte, len(prefix))
	copy(end, prefix)

	for {
		if end[len(end)-1] != byte(255) {
			end[len(end)-1]++
			break
		} else {
			end = end[:len(end)-1]
			if len(end) == 0 {
				end = nil
				break
			}
		}
	}
	return end
}

func randBytes(numBytes int) []byte {
	b := make([]byte, numBytes)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
