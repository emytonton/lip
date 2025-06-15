package pangram

import "strings"

// IsPangram checks if a string contains all letters of the English alphabet
func IsPangram(input string) bool {
	seen := make(map[rune]bool)
	
	for _, char := range strings.ToLower(input) {
		if char >= 'a' && char <= 'z' {
			seen[char] = true
		}
	}
	
	return len(seen) == 26
}