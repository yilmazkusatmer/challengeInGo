package palindrome

import (
	"fmt"
	"testing"
)

var tests = []struct {
	name  string
	input string
	want  bool
}{{fmt.Sprintf("isPalindrome('%v')", "Otto"), "Otto", true},
	{fmt.Sprintf("isPalindrome('%v')", "Affa"), "Affa", true},
	{fmt.Sprintf("isPalindrome('%v')", "Anna"), "Anna", true},
	{fmt.Sprintf("isPalindrome('%v')", "abcba"), "Otto", true},
	{fmt.Sprintf("isPalindrome('%v')", "abcabc"), "abcabc", false},
	{fmt.Sprintf("isPalindrome('%v')", "moon"), "moon", false},
	{fmt.Sprintf("isPalindrome('%v')", "123321"), "123321", true},
	{fmt.Sprintf("isPalindrome('%v')", "a"), "a", true},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.input); got != tt.want {
				t.Errorf("IsPalindrome(%v) produce %v, but expected %v", tt.input, got, tt.want)
			}
		})
	}
}
