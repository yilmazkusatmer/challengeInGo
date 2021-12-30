package duplicatechars

import (
	"strings"
)

func ContainDuplicateCharacters(input string) bool {
	duplicateMap := make(map[string]int)
	inputCharacters := []rune(strings.ToLower(input))
	for _, c := range inputCharacters {
		value := string(c)
		if count, found := duplicateMap[value]; found {
			duplicateMap[value] = count + 1
		} else {
			duplicateMap[value] = 1
		}
	}
	return len(duplicateMap) != len(input)
}
