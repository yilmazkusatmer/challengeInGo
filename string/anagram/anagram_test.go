package anagram

import "testing"

type args struct {
	inputFirst  string
	inputSecond string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"is anagram", args{"otto", "toto"}, true},
	{"is anagram", args{"mary", "army"}, true},
	{"is anagram", args{"Ananas", "Bananas"}, false},
}

func TestIsAnagram(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := IsAnagram(tt.args.inputFirst, tt.args.inputSecond); result != tt.want {
				t.Errorf("IsAnagram(%v, %v) produce %v, but expected %v", tt.args.inputFirst, tt.args.inputSecond, result, tt.want)
			}
		})
	}
}
