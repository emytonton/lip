package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && c != '\''
	}

	words := strings.FieldsFunc(s, f)

	var acronym strings.Builder

	
	for _, word := range words {
		
		firstLetter := []rune(word)[0]
		acronym.WriteRune(firstLetter)
	}


	return strings.ToUpper(acronym.String())
}
