package duplicatechars

import "testing"

var tests = []struct {
	name  string
	input string
	want  bool
}{
	{"contain duplicate characters in Otto", "Otto", true},
	{"contain duplicate characters in Ronald", "Ronald", false},
	{"contain duplicate characters in Anna", "Anna", true},
	{"contain duplicate characters in Mimimi", "Mimimi", true},
	{"contain duplicate characters in 12345", "12345", false},
}

func TestContainDuplicateCharacters(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := ContainDuplicateCharacters(tt.input); result != tt.want {
				t.Errorf("ContainDuplicateCharacters(#{tt.input}) produce #{result}, but expected #{tt.want}")
			}
		})
	}
}
