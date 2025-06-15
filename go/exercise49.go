package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)


	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}

	
	isQuestion := strings.HasSuffix(trimmedRemark, "?")
	isShouting := hasLetters(trimmedRemark) && strings.ToUpper(trimmedRemark) == trimmedRemark

	
	if isShouting && isQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if isShouting {
		return "Whoa, chill out!"
	}
	if isQuestion {
		return "Sure."
	}

	
	return "Whatever."
}

func hasLetters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
