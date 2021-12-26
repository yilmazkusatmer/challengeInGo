package perfectnumber

import (
	"reflect"
	"testing"
)

var Tests = []struct {
	name   string
	input  int
	result []int
}{
	{"calculate perfectnumber for 6", 6, []int{6}},
	{"calculate perfectnumber for 1000", 1000, []int{6, 28, 496}},
	{"calculate perfectnumber for 10000", 10000, []int{6, 28, 496, 8128}},
}

func TestCalcPerferctNumbers(t *testing.T) {
	for _, tt := range Tests {
		t.Run(tt.name, func(t *testing.T) {
			perfectNumbers := CalcPerfectNumbers(tt.input)
			if !reflect.DeepEqual(perfectNumbers, tt.result) {
				t.Errorf("CalcPerfectNumbers(%v) produce %v, but expected %v", tt.input, perfectNumbers, tt.result)
			}
		})
	}
}
