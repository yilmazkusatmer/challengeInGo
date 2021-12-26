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
	{"perfectnumber6", 6, []int{6}},
	{"perfectnumber1000", 1000, []int{6, 28, 496}},
	{"perfectnumber10000", 10000, []int{6, 28, 496, 8128}},
}

func TestCalcPerferctNumbers(t *testing.T) {
	for _, tt := range Tests {
		t.Run(tt.name, func(t *testing.T) {
			calcResult := CalcPerfectNumbers(tt.input)
			if !reflect.DeepEqual(calcResult, tt.result) {
				t.Errorf("Calc(%v) produce %v, but expected %v", 1000, calcResult, tt.result)
			}
		})
	}
}
