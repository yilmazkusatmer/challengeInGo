package palindrome

import (
	"testing"
)

var tests = []struct {
	name  string
	input string
	want  bool
}{{"check palindrome for Otto", "Otto", true},
	{"check palindrome for Affa", "Affa", true},
	{"check palindrome for Anna", "Anna", true},
	{"check palindrome for abcba", "abcba", true},
	{"check palindrom for abcabc", "abcabc", false},
	{"check palindrome for moon", "moon", false},
	{"check palindrome for 123321", "123321", true},
	{"check palindrome for a", "a", true},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := IsPalindrome(tt.input); result != tt.want {
				t.Errorf("IsPalindrome(%v) produce %v, but expected %v", tt.input, result, tt.want)
			}
		})
	}
}
