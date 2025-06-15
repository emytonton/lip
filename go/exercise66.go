package armstrong

import "math"

func IsNumber(n int) bool {
	if n < 0 {
		return false
	}

	numDigits := countDigits(n)

	sum := 0
	temp := n

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}

	return sum == n
}


func countDigits(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}