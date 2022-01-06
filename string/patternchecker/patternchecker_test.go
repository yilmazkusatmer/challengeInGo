package patternchecker

import "testing"

type args struct {
	input   string
	pattern string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"check pattern for xyxy", args{"tom tim tom tim", "xyxy"}, true},
	{"check pattern for xyyy", args{"foo fan fan foo", "xyyy"}, false},
	{"check pattern for xyyx", args{"tom tim tim tom", "xyyx"}, true},
	{"check pattern for xxxx", args{"tom tom tom tim", "xxxx"}, false},
}

func TestMatchPattern(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := MatchPattern(tt.args.input, tt.args.pattern); result != tt.want {
				t.Errorf("MatchPattern(%v, %v) produces %v, but expected %v",
					tt.args.input, tt.args.pattern, result, tt.want)
			}
		})
	}
}
