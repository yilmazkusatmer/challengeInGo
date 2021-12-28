package armstrong

import (
	"math"
)

func CalcArmstrongNumbers() []int {
	result := make([]int, 0)
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			for k := 1; k <= 10; k++ {
				value := i*100 + j*10 + k
				cubicValue := int(math.Pow(float64(i), 3) + math.Pow(float64(j), 3) + math.Pow(float64(k), 3))
				if cubicValue == value {
					result = append(result, cubicValue)
				}
			}
		}
	}
	return result
}
