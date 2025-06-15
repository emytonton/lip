package prime

import (
	"errors"
	"math"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(num)))
	for i := 3; i <= limit; i += 2 {
		if num%i == 0 {
			return false
		}
	}


	return true
}


func Nth(n int) (int, error) {
	
	if n < 1 {
		return 0, errors.New("prime is not defined for non-positive n")
	}


	primesFound := 0
	candidate := 2

	
	for {
		if isPrime(candidate) {
			primesFound++
			
			if primesFound == n {
				return candidate, nil
			}
		}
		candidate++
	}
}
