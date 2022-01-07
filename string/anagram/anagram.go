package anagram

import (
	"reflect"
	"strings"
)

func IsAnagram(inputFirst, inputSecond string) bool {
	return reflect.DeepEqual(groupBy(inputFirst), groupBy(inputSecond))
}

func groupBy(input string) map[string]int {
	characters := []rune(input)
	charactersCountMap := make(map[string]int, 0)
	for _, character := range characters {
		value := strings.ToLower(string(character))
		count, found := charactersCountMap[value]
		if found {
			charactersCountMap[value] = count + 1
		} else {
			charactersCountMap[value] = 1
		}
	}
	return charactersCountMap
}
