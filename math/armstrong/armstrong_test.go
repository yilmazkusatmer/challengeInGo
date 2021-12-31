package armstrong

import (
	"reflect"
	"testing"
)

func TestCalcArmstrongNumbers(t *testing.T) {
	t.Run("CalcArmstrongNumbers", func(t *testing.T) {
		expectedResult := []int{153, 371}
		if result := CalcArmstrongNumbers(); !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("Expected result is %v, but got %v", expectedResult, result)
		}
	})
}
