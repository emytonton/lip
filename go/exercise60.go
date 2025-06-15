package wordcount

import (
	"regexp"
	"strings"
)


type Frequency map[string]int


func WordCount(phrase string) Frequency {
	normalizedPhrase := strings.ToLower(phrase)
    
	re := regexp.MustCompile(`\b[a-z0-9]+(?:'[a-z0-9]+)*\b`)

	words := re.FindAllString(normalizedPhrase, -1)

	counts := make(Frequency)
	for _, word := range words {
		counts[word]++
	}

	return counts
}
