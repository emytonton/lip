package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {

	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	
	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}


	nums := make([]int, len(digits))
	for i, char := range digits {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
		
			return 0, errors.New("digits input must only contain digits")
		}
		nums[i] = digit
	}

	if span == 0 {
		return 1, nil
	}


	var maxProduct int64 = 0

	for i := 0; i <= len(nums)-span; i++ {
	
		series := nums[i : i+span]

	
		var currentProduct int64 = 1
		for _, digit := range series {
			currentProduct *= int64(digit)
		}

	
		if currentProduct > maxProduct {
			maxProduct = currentProduct
		}
	}

	return maxProduct, nil
}
