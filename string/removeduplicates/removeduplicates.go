package removeduplicates

import (
	"strings"
)

func RemoveDuplicates(input string) string {
	chars := []rune(input)
	nonDuplicateMap := make(map[string]string)
	result := ""
	for _, c := range chars {
		value := string(c)
		lowerCharacter := strings.ToLower(value)
		_, found := nonDuplicateMap[lowerCharacter]
		if !found {
			nonDuplicateMap[lowerCharacter] = value
			result += value
		}
	}
	return result
}
