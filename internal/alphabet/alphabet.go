package alphabet

import (
	"regexp"
	"strings"
)

var StandardAlphabet = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func GetAlphabetChar(i int, ab []rune) rune {
	for i >= 26 || i < 0 {
		if i >= 26 {
			i -= 26
		} else if i < 0 {
			i += 26
		}
	}
	return ab[i]
}

func GetCharNumber(x rune, ab []rune) int {
	for i, ac := range ab {
		if ac == x {
			return i
		}
	}
	return -1
}

var replacementRegex = regexp.MustCompile(`[^A-Z ]`)

func FilterNonAlphabet(in string) string {
	return strings.TrimSpace(replacementRegex.ReplaceAllString(in, ""))
}
