package armstrong

import (
	"reflect"
	"testing"
)

func TestCalcArmstrongNumbers(t *testing.T) {
	t.Run("CalcArmstrongNumbers", func(t *testing.T) {
		result := CalcArmstrongNumbers()
		expectedResult := []int{153, 371}
		if !reflect.DeepEqual(result, expectedResult) {
			t.Errorf("Expected result is %v, but got %v", expectedResult, result)
		}
	})
}
