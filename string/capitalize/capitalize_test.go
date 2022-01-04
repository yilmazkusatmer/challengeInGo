package capitalize

import "testing"

var tests = []struct {
	name  string
	input string
	want  string
}{
	{"Capitalize words", "this is very beautiful", "This Is Very Beautiful"},
	{"Capitalize words", "golang is cool", "Golang Is Cool"},
}

func TestCapitalize(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := Capitalize(tt.input); result != tt.want {
				t.Errorf("Capitalize('%v') produce '%v', but expected '%v'", tt.input, result, tt.want)
			}
		})
	}
}
