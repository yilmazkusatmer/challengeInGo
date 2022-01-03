package morsecode

import "testing"

var tests = []struct {
	name  string
	input string
	want  string
}{
	{"translate hello to morsecode", "hello", ".... . .-.. .-.. ---"},
	{"translate sunday to morsecode", "sunday", "... ..- -. -.. .- -.--"},
	{"translate earthquake to morsecode", "earthquake", ". .- .-. - .... --.- ..- .- -.- ."},
	{"translate Äthiopien to morsecode", "Äthiopien", ".-.- - .... .. --- .--. .. . -."},
}

func TestToMorseCode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := TranslateToMorseCode(tt.input); result != tt.want {
				t.Errorf("TranslateToMorseCode(%v) produce %v, but expected %v", tt.input, result, tt.want)
			}
		})
	}
}
