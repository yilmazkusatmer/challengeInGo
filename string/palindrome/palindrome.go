package palindrome

import (
	"strings"
)

func IsPalindrome(input string) bool {
	if strings.TrimSpace(input) == "" {
		return false
	}
	characters := []rune(strings.ToLower(input))
	for i := 0; i < len(characters)/2; i++ {
		lastElementIndex := len(characters) - 1 - i
		if characters[i] != characters[lastElementIndex] {
			return false
		}
	}
	return true
}
