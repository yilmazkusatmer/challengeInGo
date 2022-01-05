package main

import "testing"

type args struct {
	words     []string
	seperator string
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"join strings with seperator", args{[]string{"hello", "world", "message"}, "---"},
		"hello---world---message"},
	{"join strings with seperator", args{[]string{"abc", "def", "ghi"}, "..."},
		"abc...def...ghi"},
}

func TestJoin(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := Join(tt.args.words, tt.args.seperator); result != tt.want {
				t.Errorf("Join produce %v, but expected %v", result, tt.want)
			}
		})
	}
}
