package patternchecker

import (
	"regexp"
)

func MatchPattern(input string, pattern string) bool {
	words := splitByRegEx(input, "[^\\s]+")
	patternCharacters := splitByRegEx(pattern, ".{1}")

	if (len(words) == 0 && len(patternCharacters) == 0) ||
		len(words) != len(patternCharacters) {
		return false
	}

	placeholderMap := createMap(words, patternCharacters)
	for i := 0; i < len(words); i++ {
		word := words[i]
		value := patternCharacters[i]
		assignableValue := placeholderMap[value]
		if word != assignableValue {
			return false
		}
	}
	return true
}

func splitByRegEx(input string, regEx string) []string {
	compiledRegEx := regexp.MustCompile(regEx)
	return compiledRegEx.FindAllString(input, -1)
}

func createMap(words []string, patternCharacters []string) map[string]string {
	patternMap := make(map[string]string)
	for i := 0; i < len(words); i++ {
		patternMap[patternCharacters[i]] = words[i]
	}
	return patternMap
}
