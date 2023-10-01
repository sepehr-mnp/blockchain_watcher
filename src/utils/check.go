package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

func ToBase58CheckAddress(hexString string) string {
	bytes, _ := hex.DecodeString(hexString)
	checksum := calculateChecksum(bytes)
	payload := append(bytes, checksum...)
	return encodeBase58(payload)
}

func calculateChecksum(data []byte) []byte {
	hash := doubleSha256(data)
	return hash[:4]
}

func doubleSha256(data []byte) []byte {
	sha256Hash := sha256.Sum256(data)
	sha256Hash2 := sha256.Sum256(sha256Hash[:])
	return sha256Hash2[:]
}

func encodeBase58(data []byte) string {
	base58Alphabet := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	x := new(big.Int).SetBytes(data)
	base := big.NewInt(58)
	zero := big.NewInt(0)

	var result []byte
	for x.Cmp(zero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, base, mod)
		result = append([]byte{base58Alphabet[mod.Int64()]}, result...)
	}

	for _, b := range data {
		if b != 0 {
			break
		}
		result = append([]byte{base58Alphabet[0]}, result...)
	}

	return string(result)
}
