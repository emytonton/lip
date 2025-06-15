package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	
	for _, char := range strings.ToLower(word) {
		if char == ' ' || char == '-' {
			continue
		}
		if !unicode.IsLetter(char) {
			continue
		}
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}