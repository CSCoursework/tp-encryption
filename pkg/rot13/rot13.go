package rot13

import (
	"github.com/CSCoursework/tp-encryption/internal/alphabet"
	"github.com/CSCoursework/tp-encryption/pkg/caesar"
)

func CoreRot13(text string) string {
	return caesar.DoCoreCaesar(text, 13, alphabet.Standard, alphabet.Standard)
}

func Encrypt(plaintext string) string {
	return CoreRot13(plaintext)
}

func Decrypt(ciphertext string) string {
	return CoreRot13(ciphertext)
}