package keyedcaesar

import (
	"errors"
	"strings"

	"github.com/CSCoursework/tp-encryption/internal/alphabet"
	"github.com/CSCoursework/tp-encryption/pkg/caesar"
)

func makeKeyedAlphabet(key string) ([]rune, error) {
	seen := make(map[rune]bool)

	key = alphabet.FilterNonAlphabet(key)
	key = strings.ReplaceAll(key, " ", "")

	newAlphabet := make([]rune, len(alphabet.Standard))
	copy(newAlphabet, alphabet.Standard)

	for keyIdx := len(key) - 1; keyIdx >= 0; keyIdx -= 1 {
		char := rune(key[keyIdx])
		if seenBefore := seen[char]; seenBefore {
			return nil, errors.New("duplicated character in key")
		} else {
			seen[char] = true
		}

		// find character location in newAlphabet
		var location int
		for i, c := range newAlphabet {
			if c == char {
				location = i
				break
			}
		}

		newAlphabet = append(newAlphabet[:location], newAlphabet[location+1:]...)
		newAlphabet = append([]rune{char}, newAlphabet...)
	}

	return newAlphabet, nil
}

func Encrypt(plaintext, key string, shift int) (string, error) {
	ab, err := makeKeyedAlphabet(key)
	if err != nil {
		return "", err
	}
	return caesar.DoCoreCaesar(plaintext, shift, alphabet.Standard, ab), nil
}

func Decrypt(ciphertext, key string, shift int) (string, error) {
	ab, err := makeKeyedAlphabet(key)
	if err != nil {
		return "", err
	}
	return caesar.DoCoreCaesar(ciphertext, -shift, ab, alphabet.Standard), nil
}
