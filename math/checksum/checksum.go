package checksum

func CalcChecksum(num string) int {
	strValues := []rune(num)
	result := 0
	for i, n := range strValues {
		number := int(n - '0')
		result += (i + 1) * number
	}
	return result % 10
}
