package vigenere

import (
	"github.com/CSCoursework/tp-encryption/internal/alphabet"
)

func shiftAlphabetToCharacter(target rune, ab []rune) []rune {
	newAb := make([]rune, len(ab))
	copy(newAb, ab)

	for newAb[0] != target {
		last := newAb[len(newAb)-1]
		newAb = append([]rune{last}, newAb[:len(newAb)-1]...)
	}

	return newAb
}

func DoCoreVigenere(text, key string, useShiftedForOutput bool) string {
	text = alphabet.FilterNonAlphabet(text)
	key = alphabet.FilterNonAlphabet(key)

	var output []rune
	var whitespaceCount int // adding whitespace can mess up the key lookup

	for i, char := range text {

		toAppend := char

		if char != ' ' {
			keyTarget := rune(key[(i-whitespaceCount)%len(key)])

			shiftedAb := shiftAlphabetToCharacter(keyTarget, alphabet.Standard)

			if useShiftedForOutput {
				toAppend = shiftedAb[alphabet.GetCharNumber(char, alphabet.Standard)]
			} else {
				toAppend = alphabet.Standard[alphabet.GetCharNumber(char, shiftedAb)]
			}

		} else {
			whitespaceCount += 1
		}

		output = append(output, toAppend)

	}

	return string(output)
}

func Encrypt(plaintext, key string) string {
	return DoCoreVigenere(plaintext, key, true)
}

func Decrypt(ciphertext, key string) string {
	return DoCoreVigenere(ciphertext, key, false)
}
