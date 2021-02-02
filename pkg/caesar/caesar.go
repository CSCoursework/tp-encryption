package caesar

import (
	"regexp"
	"strings"

	"github.com/CSCoursework/tp-encryption/internal/alphabet"
)


func DoCoreCaesar(text string, modifier int, ab []rune) string {
	// make upper case and filter out non-space and non-alphanumeric characters
	text = strings.ToUpper(text)
	text = alphabet.FilterNonAlphabet(text)

	var o []rune
	for _, char := range text {
		if char != ' ' {
			cn := alphabet.GetCharNumber(char, ab)
			o = append(o, alphabet.GetAlphabetChar(cn+modifier, ab))
		} else {
			o = append(o, ' ')
		}
	}
	return string(o)
}

func Encrypt(plaintext string, shift int) string {
	return DoCoreCaesar(plaintext, shift, alphabet.StandardAlphabet)
}

func Decrypt(ciphertext string, shift int) string {
	return DoCoreCaesar(ciphertext, -shift, alphabet.StandardAlphabet)
}
