package isbn

import (
	"strings"
	"unicode"
)


func IsValidISBN(isbn string) bool {
	
	cleaned := strings.ReplaceAll(isbn, "-", "")

	
	if len(cleaned) != 10 {
		return false
	}

	sum := 0
	for i, c := range cleaned {
		var value int
		if i == 9 && (c == 'X' || c == 'x') {
			value = 10
		} else if unicode.IsDigit(c) {
			value = int(c - '0')
		} else {
			return false 
		}
		sum += value * (10 - i)
	}

	return sum%11 == 0
}