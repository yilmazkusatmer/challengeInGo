package removeduplicates

import "testing"

var tests = []struct {
	name  string
	input string
	want  string
}{
	{"Remove duplicate from Otto", "Otto", "Ot"},
	{"Remove duplicate from abcba", "abcba", "abc"},
	{"Remove duplicate from moon", "moon", "mon"},
	{"Remove duplicate from michael", "michael", "michaela"},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := RemoveDuplicates(tt.input); result != tt.want {
				t.Errorf("RemoveDuplicates produce %v, but expected %v", result, tt.want)
			}
		})
	}
}
