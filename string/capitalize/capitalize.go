package capitalize

import (
	"regexp"
	"strings"
)

func Capitalize(input string) string {
	splitBySpace := regexp.MustCompile("\\w+")
	return splitBySpace.ReplaceAllStringFunc(input, func(word string) string {
		return strings.Title(word)
	})
}
