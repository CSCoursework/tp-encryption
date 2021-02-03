package xor

import (
	"encoding/hex"
	"errors"
)

func DoCoreXor(inp []byte, key int) ([]byte, error) {
	if key < 0 || key > 255 {
		return nil, errors.New("key out of range - must be 0 <= n <= 255")
	}

	byteKey := byte(key)

	var o []byte
	for _, char := range inp {
		o = append(o, char^byteKey)
	}

	return o, nil
}

func Encrypt(plaintext string, key int) (string, error) {
	bx, err := DoCoreXor([]byte(plaintext), key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bx), nil
}

func Decrypt(ciphertext string, key int) (string, error) {
	hx, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	pt, err := DoCoreXor(hx, key)
	return string(pt), err
}
