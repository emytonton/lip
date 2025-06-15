package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	
	var intermediate strings.Builder


	for _, r := range s {
		
		if unicode.IsLetter(r) {
			
			lower := unicode.ToLower(r)
			
			encodedChar := 'z' - lower + 'a'
			intermediate.WriteRune(encodedChar)
			continue
		}

		
		if unicode.IsDigit(r) {
			intermediate.WriteRune(r)
		}
		
	}

	
	var result strings.Builder
	unformatted := intermediate.String()
	for i, r := range unformatted {
		
		if i > 0 && i%5 == 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(r)
	}

	return result.String()
}
