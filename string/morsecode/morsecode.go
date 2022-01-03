package morsecode

import (
	"strings"
)

var morseCodeMap = map[string]string{
	"A": ".-", "B": "-...", "C": "-.-.", "D": "-..",
	"E": ".", "F": "..-.", "G": "..-.", "H": "....",
	"I": "..", "J": ".---", "K": "-.-", "L": ".-..",
	"M": "--", "N": "-.", "O": "---", "P": ".--.",
	"Q": "--.-", "R": ".-.", "S": "...", "T": "-",
	"U": "..-", "V": "...-", "W": ".--", "X": "-..-",
	"Y": "-.--", "Z": "--..", "Ä": ".-.-", "Ö": "---.",
	"ß": "......", "Ü": "..--",
}

func TranslateToMorseCode(input string) string {
	characters := []rune(input)
	result := ""
	for index, character := range characters {
		value := string(character)
		if index < len(characters)-1 {
			result += morseCodeMap[strings.ToUpper(value)] + " "
		} else {
			result += morseCodeMap[strings.ToUpper(value)]
		}
	}
	return result
}
