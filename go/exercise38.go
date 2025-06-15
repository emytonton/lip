package luhn

import (
	"strings"
	"unicode"
)


func Valid(id string) bool {

	stripped := strings.ReplaceAll(id, " ", "")

	
	if len(stripped) <= 1 {
		return false
	}

	var sum int
	shouldDouble := len(stripped)%2 == 0 

	for _, r := range stripped {
		if !unicode.IsDigit(r) {
			return false 
		}

		digit := int(r - '0')

		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shouldDouble = !shouldDouble 
	}

	return sum%10 == 0
}