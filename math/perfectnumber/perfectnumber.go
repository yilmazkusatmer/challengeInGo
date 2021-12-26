package perfectnumber

func isPerfectNumber(num int) bool {
	sumOfResult := 1
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			sumOfResult += i
		}
	}
	return num == sumOfResult
}

func calcPerferctNumbers(num int) []int {
	calcResult := make([]int, 0)
	for i := 2; i <= num; i++ {
		if isPerfectNumber(i) {
			calcResult = append(calcResult, i)
		}
	}
	return calcResult
}
