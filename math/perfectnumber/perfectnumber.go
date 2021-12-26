package perfectnumber

func isPerfectNumber(num int) bool {
	var sumOfPerfectNumbers int = 1
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			sumOfPerfectNumbers += i
		}
	}
	return sumOfPerfectNumbers == num
}

func CalcPerfectNumbers(num int) []int {
	perfectNumbers := make([]int, 0)
	for i := 2; i <= num; i++ {
		if isPerfectNumber(i) {
			perfectNumbers = append(perfectNumbers, i)
		}
	}
	return perfectNumbers
}
