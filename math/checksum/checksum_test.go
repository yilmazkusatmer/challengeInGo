package checksum

import "testing"

var Inputs = []struct {
	name   string
	input  string
	result int
}{
	{"checksum of 1111", "1111", 0},
	{"checksum of 1234", "1234", 0},
	{"checksum of 2299", "2299", 9},
	{"checksum of 5180", "5180", 1},
	{"checksum of 876543210", "876543210", 0},
}

func TestCalcChecksum(t *testing.T) {
	for _, tt := range Inputs {
		t.Run(tt.name, func(t *testing.T) {
			if checksum := CalcChecksum(tt.input); checksum != tt.result {
				t.Errorf("CalcChecksum(#{tt.input}) produce #{checksum}, but expected #{tt.result}")
			}
		})
	}
}
